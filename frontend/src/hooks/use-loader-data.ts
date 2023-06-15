import { useLoaderData as useData } from "react-router-dom";

export const useLoaderData = <T extends (...args: any) => Promise<any> | any>() => useData() as Awaited<ReturnType<T>>;
