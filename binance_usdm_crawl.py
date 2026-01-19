#!/usr/bin/env python3
# 此脚本的作用是爬取币安衍生品交易，U本位合约下所有的技术文档。
import asyncio
import html as html_lib
import json
import re
from pathlib import Path
from urllib.parse import urljoin

from crawl4ai import AsyncWebCrawler, BrowserConfig, CrawlerRunConfig, CacheMode

BASE_URL = "https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info"
SECTION_PREFIX = "https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/"
OUTPUT_DIR = Path("binance_usdm_sidebar_pages")
SIDEBAR_WAIT_JS = "js:() => { const m = document.querySelector('.theme-doc-sidebar-menu'); return m && m.querySelectorAll('a.menu__link').length > 2; }"
SIDEBAR_READY_JS = "js:() => !!document.querySelector('#usdm-links')"

EXPAND_JS = r"""
(async () => {
  const targetPath = "/docs/zh-CN/derivatives/usds-margined-futures/";
  const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
  const doClick = (el) => {
    try {
      el.dispatchEvent(new MouseEvent("click", { bubbles: true, cancelable: true }));
    } catch (e) {
      try { el.click(); } catch (e2) {}
    }
  };
  const clickAll = (selector) => {
    document.querySelectorAll(selector).forEach((el) => {
      doClick(el);
    });
  };

  // Expand native details elements
  document.querySelectorAll('details').forEach((d) => { d.open = true; });

  // Expand common accordion toggles
  clickAll('[aria-expanded="false"]');
  clickAll('[data-state="closed"]');
  clickAll('button[aria-expanded="false"]');
  clickAll('[role="button"][aria-expanded="false"]');

  // Scroll to trigger lazy content
  for (let i = 0; i < 6; i++) {
    window.scrollTo(0, document.body.scrollHeight);
    await sleep(400);
  }

  // Wait for sidebar to render (open mobile sidebar if needed)
  let menu = null;
  for (let i = 0; i < 40; i++) {
    menu = document.querySelector(".theme-doc-sidebar-menu");
    if (!menu) {
      const toggle = document.querySelector("button.navbar__toggle");
      if (toggle) {
        doClick(toggle);
        await sleep(400);
      }
      menu = document.querySelector(".theme-doc-sidebar-menu");
    }
    if (menu && menu.querySelectorAll("a.menu__link").length > 2) break;
    await sleep(250);
  }
  if (!menu) return;

  const uCat = Array.from(menu.querySelectorAll("li.theme-doc-sidebar-item-category")).find((li) => {
    const a = li.querySelector("a.menu__link--sublist");
    if (!a) return false;
    const href = a.getAttribute("href") || "";
    const text = (a.textContent || "").trim();
    return text.includes("U本位合约") || href.includes(targetPath);
  });
  if (!uCat) return;

  // Expand all sublists inside U本位合约
  for (let i = 0; i < 30; i++) {
    let changed = false;
    const toggles = Array.from(uCat.querySelectorAll("a.menu__link--sublist"));
    for (const a of toggles) {
      if (a.getAttribute("aria-expanded") === "false") {
        doClick(a);
        changed = true;
        await sleep(120);
      }
    }
    if (!changed) break;
    await sleep(200);
  }

  const items = Array.from(uCat.querySelectorAll("li.theme-doc-sidebar-item-link a.menu__link"))
    .map((a) => {
      const href = a.getAttribute("href") || "";
      const text = (a.textContent || "").trim();
      const path = [];
      let current = a.closest("li");
      while (current) {
        const parentCat = current.closest("li.theme-doc-sidebar-item-category");
        if (!parentCat || !uCat.contains(parentCat)) break;
        const labelEl =
          parentCat.querySelector(":scope > .menu__list-item-collapsible > a.menu__link--sublist") ||
          parentCat.querySelector(":scope > a.menu__link--sublist");
        if (labelEl) {
          const label = (labelEl.textContent || "").trim();
          if (label && label !== "U本位合约") {
            path.unshift(label);
          }
        }
        current = parentCat.parentElement;
      }
      return { href, text, path };
    })
    .filter((x) => x.href && (x.href.startsWith(targetPath) || x.href.includes(targetPath)))
    .map((x) => ({ href: new URL(x.href, location.origin).href, text: x.text, path: x.path }));

  const unique = [];
  const seen = new Set();
  for (const item of items) {
    const abs = new URL(item.href, location.origin).href;
    const key = abs + "|" + item.text + "|" + (item.path || []).join(">");
    if (!seen.has(key)) {
      seen.add(key);
      unique.push({ href: abs, text: item.text, path: item.path || [] });
    }
  }
  const holder = document.createElement("pre");
  holder.id = "usdm-links";
  holder.style.display = "none";
  holder.textContent = JSON.stringify(unique);
  document.body.appendChild(holder);
})();
"""


def extract_usdm_items(html: str) -> list[dict]:
    m = re.search(r'<pre id="usdm-links"[^>]*>(.*?)</pre>', html, re.S)
    if m:
        try:
            items = m.group(1).strip()
            data = json.loads(items)
            out = []
            for item in data:
                href = item.get("href", "")
                if href.startswith(SECTION_PREFIX):
                    out.append(
                        {
                            "href": href.split("#", 1)[0],
                            "text": item.get("text", "").strip(),
                            "path": item.get("path") or [],
                        }
                    )
            return out
        except Exception:
            pass

    # Fallback: parse sidebar links directly from HTML
    items = []
    pattern = re.compile(
        r'<li class="[^"]*theme-doc-sidebar-item-link[^"]*theme-doc-sidebar-item-link-level-(\d+)[^"]*"[^>]*>\\s*'
        r'<a[^>]*class="[^"]*menu__link[^"]*"[^>]*href="([^"]+)"[^>]*>(.*?)</a>',
        re.S,
    )
    for level_s, href, text in pattern.findall(html):
        if not (href.startswith("/") or href.startswith("http")):
            continue
        clean_text = re.sub(r"<[^>]+>", "", text)
        clean_text = html_lib.unescape(clean_text).strip()
        items.append(
            {
                "href": href,
                "text": clean_text,
                "level": int(level_s),
            }
        )
    out = []
    for item in items:
        href = item.get("href", "")
        if href.startswith("/"):
            href = "https://developers.binance.com" + href
        if href.startswith(SECTION_PREFIX):
            out.append(
                {
                    "href": href.split("#", 1)[0],
                    "text": item.get("text", "").strip(),
                    "path": [],
                }
            )
    return out


def extract_page_title(html: str) -> str:
    m = re.search(r"<h1[^>]*>(.*?)</h1>", html, re.S)
    if m:
        text = re.sub(r"<[^>]+>", "", m.group(1)).strip()
        if text:
            return text
    m = re.search(r"<title[^>]*>(.*?)</title>", html, re.S)
    if m:
        text = re.sub(r"<[^>]+>", "", m.group(1)).strip()
        if text:
            return text
    return ""


def normalize_url(href: str) -> str:
    if href.startswith("//"):
        href = "https:" + href
    return urljoin("https://developers.binance.com", href)


def safe_filename(title: str, url: str) -> str:
    name = re.sub(r"[\\/:*?\"<>|]", "", title).strip()
    name = re.sub(r"\\s+", "_", name)
    if not name:
        slug = url.replace(SECTION_PREFIX, "").strip("/") or "index"
        slug = re.sub(r"[^a-zA-Z0-9_-]+", "_", slug)
        name = slug
    return f"{name}.md"


def safe_dirname(title: str) -> str:
    name = re.sub(r"[\\/:*?\"<>|]", "", title).strip()
    name = re.sub(r"\\s+", "_", name)
    return name or "未命名"


def unique_path(path: Path, url: str) -> Path:
    if not path.exists():
        return path
    slug = url.replace(SECTION_PREFIX, "").strip("/") or "page"
    slug = re.sub(r"[^a-zA-Z0-9_-]+", "_", slug)
    return path.with_name(f"{path.stem}_{slug}{path.suffix}")


def get_markdown(result) -> str:
    md = getattr(result, "markdown", "")
    if isinstance(md, str):
        return md
    raw = getattr(md, "raw_markdown", "")
    if raw:
        return raw
    return getattr(md, "fit_markdown", "") or ""


async def crawl_page(
    crawler: AsyncWebCrawler,
    url: str,
    *,
    js_code=None,
    wait_for=None,
    js_only: bool = False,
    session_id=None,
    delay: float = 0.5,
):
    run_cfg = CrawlerRunConfig(
        cache_mode=CacheMode.BYPASS,
        js_code=js_code,
        wait_for=wait_for,
        wait_until="networkidle",
        page_timeout=120000,
        delay_before_return_html=delay,
        js_only=js_only,
        session_id=session_id,
    )
    result = await crawler.arun(url=url, config=run_cfg)
    return result


async def extract_sidebar_items(crawler: AsyncWebCrawler) -> list[dict]:
    session_id = "usdm_sidebar_session"

    base = await crawl_page(
        crawler,
        BASE_URL,
        wait_for=SIDEBAR_WAIT_JS,
        session_id=session_id,
        delay=1.0,
    )
    if not base.success:
        raise RuntimeError(f"Base page crawl failed: {base.error_message}")

    expanded = await crawl_page(
        crawler,
        BASE_URL,
        js_code=[EXPAND_JS],
        wait_for=SIDEBAR_READY_JS,
        js_only=True,
        session_id=session_id,
        delay=1.0,
    )
    if not expanded.success:
        raise RuntimeError(f"Sidebar expand failed: {expanded.error_message}")

    html = expanded.html or ""
    items = extract_usdm_items(html)
    if not items:
        has_menu = ".theme-doc-sidebar-menu" in html
        print(f"[WARN] Sidebar links not found in HTML; has_sidebar={has_menu}")

    # Cleanup session
    try:
        await crawler.crawler_strategy.kill_session(session_id)
    except Exception:
        pass

    merged = {}
    for item in items:
        href = item.get("href", "")
        title = item.get("text", "")
        path = item.get("path") or []
        if not href:
            continue
        score = (len(path), len(title))
        if href not in merged or score > merged[href]["score"]:
            merged[href] = {"text": title, "path": path, "score": score}
    return [
        {"href": url, "text": data["text"], "path": data["path"]}
        for url, data in sorted(merged.items())
    ]


async def main():
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)

    browser_cfg = BrowserConfig(
        headless=True,
        java_script_enabled=True,
        viewport_width=1440,
        viewport_height=900,
        user_agent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
        verbose=False,
    )

    async with AsyncWebCrawler(config=browser_cfg) as crawler:
        items = await extract_sidebar_items(crawler)
        print(f"[INFO] Found {len(items)} USDM menu links from sidebar")

        for item in items:
            url = item["href"]
            title = item.get("text", "").strip()
            path_parts = item.get("path") or []
            result = await crawl_page(crawler, url, wait_for="css:main", delay=0.3)
            if not result.success:
                print(f"[WARN] Crawl failed: {url} -> {result.error_message}")
                continue
            md = get_markdown(result)
            if not title:
                title = extract_page_title(result.html or "")
            dir_path = OUTPUT_DIR
            for part in path_parts:
                dir_path = dir_path / safe_dirname(part)
            dir_path.mkdir(parents=True, exist_ok=True)
            out_path = dir_path / safe_filename(title, url)
            out_path = unique_path(out_path, url)
            out_path.write_text(md, encoding="utf-8")
            print(f"[OK] {url} -> {out_path}")


if __name__ == "__main__":
    asyncio.run(main())
