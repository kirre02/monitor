import { CheckApi, SiteApi, StatusApi } from "monitor-sdk";
import { Configuration } from "monitor-sdk/runtime";

export interface APIBridge{
    checkApi: CheckApi
    siteApi: SiteApi
    statusApi: StatusApi 
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
    const statusApi = new StatusApi(config)

    apiBridgeCache = {
      siteApi,
      checkApi,
      statusApi
    }
    
    return apiBridgeCache;
}
