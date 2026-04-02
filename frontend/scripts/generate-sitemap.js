import { writeFileSync } from "fs";

const BASE_URL = "https://your-domain.com"; // change this

// Public routes only
const routes = [
  "",
  "/about",
  "/terms",
  "/privacy",
];

function generateSitemap() {
  const now = new Date().toISOString();

  const urls = routes
    .map(
      (route) => `
  <url>
    <loc>${BASE_URL}${route}</loc>
    <lastmod>${now}</lastmod>
    <changefreq>monthly</changefreq>
    <priority>${route === "" ? "1.0" : "0.6"}</priority>
  </url>`
    )
    .join("");

  return `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${urls}
</urlset>`;
}

writeFileSync("static/sitemap.xml", generateSitemap());

console.log("Sitemap generated.");