import { useEffect, useState } from "react";

type FetchOptions = {
  autoFetch?: boolean; // por defecto true
  params?: string; // por ejemplo un ID
};

type FetchState<T> = {
  data: T | null;
  loading: boolean;
  error: string | null;
  fetchData: () => Promise<void>; // manual fetch
};

export function useFetch<T = unknown>(
  baseUrl: string,
  options: FetchOptions = {}
): FetchState<T> {
  const { autoFetch = true, params = "" } = options;

  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(autoFetch);
  const [error, setError] = useState<string | null>(null);

  const fullurl = params ? `${baseUrl}/${params}` : baseUrl;

  const fetchData = async () => {
    try {
      setLoading(true);
      setError(null);

      const response = await fetch(fullurl);

      if (!response.ok)
        throw new Error(`Error en la solicitud ${response.status}`);

      const jsonResponse = await response.json();
      setData(jsonResponse.data || jsonResponse);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (autoFetch) fetchData();
  }, [fullurl]);

  return { data, loading, error, fetchData };
}
