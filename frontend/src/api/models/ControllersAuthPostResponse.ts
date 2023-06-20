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
import type { ControllersAuthTokens } from "./ControllersAuthTokens";
import {
	ControllersAuthTokensFromJSON,
	ControllersAuthTokensFromJSONTyped,
	ControllersAuthTokensToJSON,
} from "./ControllersAuthTokens";
import type { ModelsUser } from "./ModelsUser";
import { ModelsUserFromJSON, ModelsUserFromJSONTyped, ModelsUserToJSON } from "./ModelsUser";

/**
 * @export
 * @interface ControllersAuthPostResponse
 */
export interface ControllersAuthPostResponse {
	/**
	 * @memberof ControllersAuthPostResponse
	 * @type {ControllersAuthTokens}
	 */
	tokens: ControllersAuthTokens;
	/**
	 * @memberof ControllersAuthPostResponse
	 * @type {ModelsUser}
	 */
	user: ModelsUser;
}

/** Check if a given object implements the ControllersAuthPostResponse interface. */
export function instanceOfControllersAuthPostResponse(value: object): boolean {
	let isInstance = true;
	isInstance = isInstance && "tokens" in value;
	isInstance = isInstance && "user" in value;

	return isInstance;
}

export function ControllersAuthPostResponseFromJSON(json: any): ControllersAuthPostResponse {
	return ControllersAuthPostResponseFromJSONTyped(json, false);
}

export function ControllersAuthPostResponseFromJSONTyped(
	json: any,
	ignoreDiscriminator: boolean,
): ControllersAuthPostResponse {
	if (json === undefined || json === null) {
		return json;
	}
	return {
		tokens: ControllersAuthTokensFromJSON(json["tokens"]),
		user: ModelsUserFromJSON(json["user"]),
	};
}

export function ControllersAuthPostResponseToJSON(value?: ControllersAuthPostResponse | null): any {
	if (value === undefined) {
		return undefined;
	}
	if (value === null) {
		return null;
	}
	return {
		tokens: ControllersAuthTokensToJSON(value.tokens),
		user: ModelsUserToJSON(value.user),
	};
}