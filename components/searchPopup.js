import styles from "../styles/searchPopup.module.css";
import { ListGroup, ListGroupItem } from "react-bootstrap";

export default function SearchPopup(props) {
  return (
    <ListGroup
      className={`${styles.searchPopup} shadow rounded bg-body ${
        props.hidden ? styles.hidden : ""
      }`}
    >
      {props.items.map((item, index) => (
        <button
          type="button"
          key={index}
          tabIndex={index + 2}
          onClick={() => props.onSelect(item)}
          className="list-group-item list-group-item-action"
        >
          {item}
        </button>
      ))}
    </ListGroup>
  );
}
