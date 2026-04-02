export const prerender = false;
export const ssr = false;

import { siteConfig } from "$lib/config";

export function load() {
  return {
    meta: {
      title: siteConfig.siteName,
      description: siteConfig.siteDescription,
      path: "",
    },
  };
}
