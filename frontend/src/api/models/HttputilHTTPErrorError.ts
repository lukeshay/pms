/* tslint:disable */
/* eslint-disable */
/**
 * Some API
 * This is a sample server celler server.
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface HttputilHTTPErrorError
 */
export interface HttputilHTTPErrorError {
    /**
     * 
     * @type {string}
     * @memberof HttputilHTTPErrorError
     */
    message?: string;
}

/**
 * Check if a given object implements the HttputilHTTPErrorError interface.
 */
export function instanceOfHttputilHTTPErrorError(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HttputilHTTPErrorErrorFromJSON(json: any): HttputilHTTPErrorError {
    return HttputilHTTPErrorErrorFromJSONTyped(json, false);
}

export function HttputilHTTPErrorErrorFromJSONTyped(json: any, ignoreDiscriminator: boolean): HttputilHTTPErrorError {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'message': !exists(json, 'message') ? undefined : json['message'],
    };
}

export function HttputilHTTPErrorErrorToJSON(value?: HttputilHTTPErrorError | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'message': value.message,
    };
}
