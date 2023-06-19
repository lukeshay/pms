import { useLoaderData as useData } from "react-router-dom";

export const useLoaderData = <T extends (...args: any) => Promise<any>>() => useData() as Awaited<ReturnType<T>>;
