/* tslint:disable */
/* eslint-disable */
/**
 * Some API This is a sample server celler server.
 *
 * The version of the OpenAPI document: 1.0 Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech Do not edit the class manually.
 */

import { exists, mapValues } from "../runtime";
/**
 * @export
 * @interface ControllersBooksDeleteResponse
 */
export interface ControllersBooksDeleteResponse {
	/**
	 * @memberof ControllersBooksDeleteResponse
	 * @type {string}
	 */
	bookId: string;
}

/** Check if a given object implements the ControllersBooksDeleteResponse interface. */
export function instanceOfControllersBooksDeleteResponse(value: object): boolean {
	let isInstance = true;
	isInstance = isInstance && "bookId" in value;

	return isInstance;
}

export function ControllersBooksDeleteResponseFromJSON(json: any): ControllersBooksDeleteResponse {
	return ControllersBooksDeleteResponseFromJSONTyped(json, false);
}

export function ControllersBooksDeleteResponseFromJSONTyped(
	json: any,
	ignoreDiscriminator: boolean,
): ControllersBooksDeleteResponse {
	if (json === undefined || json === null) {
		return json;
	}
	return {
		bookId: json["bookId"],
	};
}

export function ControllersBooksDeleteResponseToJSON(value?: ControllersBooksDeleteResponse | null): any {
	if (value === undefined) {
		return undefined;
	}
	if (value === null) {
		return null;
	}
	return {
		bookId: value.bookId,
	};
}
