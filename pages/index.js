import Head from "next/head";
import styles from "../styles/Home.module.css";

import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import { Search, GitHub } from "react-feather";

import images from "../images.json";
import { useState } from "react";
import { InputGroup, Modal, Badge } from "react-bootstrap";
import SearchPopup from "../components/searchPopup";

import ReactMarkdown from "react-markdown";

function findImages(search) {
  if (search == "") {
    return [];
  }

  return images.filter((i) => i.name.includes(search));
}

export default function Home() {
  const [search, setSearch] = useState("");
  const [modalShow, setModalShow] = useState(false);

  const [image, setImage] = useState({});

  return (
    <div className={styles.container}>
      <Head>
        <title>dockerenv</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className="mb-3">dockerenv</h1>
        <Form inline style={{ position: "relative" }}>
          <InputGroup>
            <InputGroup.Prepend>
              <InputGroup.Text>
                <Search size={16} />
              </InputGroup.Text>
            </InputGroup.Prepend>
            <Form.Control
              type="text"
              name="q"
              placeholder="Find an image"
              autoFocus={true}
              autoComplete="off"
              onInput={(e) => setSearch(e.target.value)}
              tabIndex={1}
            />
          </InputGroup>

          <SearchPopup
            hidden={search == ""}
            items={findImages(search).map((i) => i.name)}
            onSelect={(item) => {
              setImage(images.find((i) => i.name == item));
              setModalShow(true);
            }}
          />
        </Form>

        <Modal size="lg" show={modalShow} onHide={() => setModalShow(false)}>
          <Modal.Header closeButton>
            <Modal.Title>{image.name}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            {image.variables &&
              Object.entries(image.variables).map(([name, variable], index) => (
                <div className="mb-2" key={name}>
                  <code>{name}</code>{" "}
                  {variable.required && (
                    <Badge variant="danger">Required</Badge>
                  )}
                  <span>
                    <ReactMarkdown>{variable.short}</ReactMarkdown>
                  </span>
                </div>
              ))}
          </Modal.Body>
        </Modal>
      </main>

      <a href="https://github.com/cjdenio/dockerenv">
        <div style={{ position: "absolute", left: 10, bottom: 10 }}>
          <GitHub size={20} />
        </div>
      </a>
    </div>
  );
}
