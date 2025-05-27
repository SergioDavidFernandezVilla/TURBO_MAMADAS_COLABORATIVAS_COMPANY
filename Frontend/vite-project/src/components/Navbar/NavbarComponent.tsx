// CSS MODULOS
import styles from "../Navbar/css/NavbarComponent.module.css";

// ICONS
import { HiBars3 } from "react-icons/hi2";

export const NavbarComponent = () => {
  return (
    <header className={styles.header__page__}>
      <nav className={styles.navbar__header__page}>
        <ul className={styles.ul__navbar__header__page}>
          <li className={styles.li__ul__navbar__header__page}>
            <a className={styles.a__li__navbar__header__page}>
              <HiBars3 className={styles.icon__a__navbar__header__page} />

              <h1 className={styles.h1__a__navbar__header__page}>
                TiendaMakaca
              </h1>
            </a>
          </li>
        </ul>
      </nav>
    </header>
  );
};
