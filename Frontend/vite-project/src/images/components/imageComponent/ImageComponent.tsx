// FUNCION UTILS
import { getImageUrl } from "../../../utils/functionsImages/GetImageUrl";

type ImgType = {
  url: string;
  alt?: string;
};

export const ImageComponent: React.FC<ImgType> = ({ url, alt }) => {
  return <img src={getImageUrl(url)} alt={alt} />;
};
