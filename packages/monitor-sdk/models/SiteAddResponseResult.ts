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
 * @interface SiteAddResponseResult
 */
export interface SiteAddResponseResult {
    /**
     * 
     * @type {Site}
     * @memberof SiteAddResponseResult
     */
    site: Site;
}

/**
 * Check if a given object implements the SiteAddResponseResult interface.
 */
export function instanceOfSiteAddResponseResult(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "site" in value;

    return isInstance;
}

export function SiteAddResponseResultFromJSON(json: any): SiteAddResponseResult {
    return SiteAddResponseResultFromJSONTyped(json, false);
}

export function SiteAddResponseResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): SiteAddResponseResult {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'site': SiteFromJSON(json['site']),
    };
}

export function SiteAddResponseResultToJSON(value?: SiteAddResponseResult | null): any {
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

