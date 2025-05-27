import { API_IMAGES } from "../API/key_API";

export const fetchImagesByProductId = async (id: number) => {
  try {
    const res = await fetch(`${API_IMAGES}/${id}`);
    const json = await res.json();
    return json.data || [];
  } catch {
    return [];
  }
};
