import { CheckApi, SiteApi } from "monitor-sdk";
import { Configuration } from "monitor-sdk/runtime";

export interface APIBridge{
    checkApi: CheckApi
    siteApi: SiteApi 
}

let apiBridgeCache: null | APIBridge = null

export function createAPIBridge() {
    if (apiBridgeCache) {
        console.log('Using cached API bridge');
        return apiBridgeCache;
    }

    const config = new Configuration({});
    
    const siteApi = new SiteApi(config)
    const checkApi = new CheckApi(config)

    apiBridgeCache = {
      siteApi,
      checkApi
    }
    
    return apiBridgeCache;
}
