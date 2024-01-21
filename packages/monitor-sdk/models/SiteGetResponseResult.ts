/* tslint:disable */
/* eslint-disable */
/**
 * Monitor Proxy API
 * Basic API proxy for Monitor
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { Site } from './Site';
import {
    SiteFromJSON,
    SiteFromJSONTyped,
    SiteToJSON,
} from './Site';

/**
 * 
 * @export
 * @interface SiteGetResponseResult
 */
export interface SiteGetResponseResult {
    /**
     * 
     * @type {Site}
     * @memberof SiteGetResponseResult
     */
    site?: Site;
}

/**
 * Check if a given object implements the SiteGetResponseResult interface.
 */
export function instanceOfSiteGetResponseResult(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SiteGetResponseResultFromJSON(json: any): SiteGetResponseResult {
    return SiteGetResponseResultFromJSONTyped(json, false);
}

export function SiteGetResponseResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): SiteGetResponseResult {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'site': !exists(json, 'site') ? undefined : SiteFromJSON(json['site']),
    };
}

export function SiteGetResponseResultToJSON(value?: SiteGetResponseResult | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'site': SiteToJSON(value.site),
    };
}

