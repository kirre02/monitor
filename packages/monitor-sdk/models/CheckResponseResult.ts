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
import type { Check } from './Check';
import {
    CheckFromJSON,
    CheckFromJSONTyped,
    CheckToJSON,
} from './Check';

/**
 * 
 * @export
 * @interface CheckResponseResult
 */
export interface CheckResponseResult {
    /**
     * 
     * @type {Check}
     * @memberof CheckResponseResult
     */
    check?: Check;
}

/**
 * Check if a given object implements the CheckResponseResult interface.
 */
export function instanceOfCheckResponseResult(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CheckResponseResultFromJSON(json: any): CheckResponseResult {
    return CheckResponseResultFromJSONTyped(json, false);
}

export function CheckResponseResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): CheckResponseResult {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'check': !exists(json, 'check') ? undefined : CheckFromJSON(json['check']),
    };
}

export function CheckResponseResultToJSON(value?: CheckResponseResult | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'check': CheckToJSON(value.check),
    };
}

