import { type Admin, type ClientResponseError, type ListResult, type Record } from "pocketbase";

/**
 * Pocketbase errors (wrapped by them)
 *
 * @example
 * const error: PocketbaseError = {
 *   url: "http://localhost:8090/api/collections/accounts/auth-with-password",
 *   status: 400,
 *   response: { code: 400, message: "Failed to authenticate.", data: {} },
 *   isAbort: false,
 *   originalError: null,
 *   name: "ClientResponseError 400",
 * };
 */
export type PocketbaseError = ClientResponseError;

export type AuthUser = Record | Admin | null;

/** Paginated list */
export type PaginatedList = ListResult;

export interface Pagination {
  page: number;
  perPage: number;
}
