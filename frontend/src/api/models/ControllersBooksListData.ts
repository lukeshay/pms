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
import type { ModelsBook } from './ModelsBook';
import {
    ModelsBookFromJSON,
    ModelsBookFromJSONTyped,
    ModelsBookToJSON,
} from './ModelsBook';

/**
 * 
 * @export
 * @interface ControllersBooksListData
 */
export interface ControllersBooksListData {
    /**
     * 
     * @type {Array<ModelsBook>}
     * @memberof ControllersBooksListData
     */
    books: Array<ModelsBook>;
}

/**
 * Check if a given object implements the ControllersBooksListData interface.
 */
export function instanceOfControllersBooksListData(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "books" in value;

    return isInstance;
}

export function ControllersBooksListDataFromJSON(json: any): ControllersBooksListData {
    return ControllersBooksListDataFromJSONTyped(json, false);
}

export function ControllersBooksListDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBooksListData {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'books': ((json['books'] as Array<any>).map(ModelsBookFromJSON)),
    };
}

export function ControllersBooksListDataToJSON(value?: ControllersBooksListData | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'books': ((value.books as Array<any>).map(ModelsBookToJSON)),
    };
}
