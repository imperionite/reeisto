export const siteConfig = {
  siteName: "REEISTO",
  siteDescription:
    "REEISTO is an academic proof-of-concept system simulating inventory and operational workflows for rare earth materials trading environments.",
  // Use Firebase final hosting URL in prod
  siteUrl: import.meta.env.VITE_SITE_URL || "http://localhost:5173",
  // backend base URL http://127.0.0.1:8080/api/v1 (dev)
  apiBaseUrl: import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1",
};
