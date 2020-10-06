import styled from "@emotion/styled";
import { Button, Form, Input, Row, Col, message } from "antd";
import { CopyOutlined } from "@ant-design/icons";
import Navbar from "components/Navbar";
import { useClipboard } from "use-clipboard-copy";
import { useEffect, useState } from "react";
import { resolve } from "path";

const StyledInput = styled(Input)`
  border-radius: 0;
  padding: 0.25rem 0.75rem;
  font-family: "Montserrat", sans-serif;
`;
const PageContainer = styled.div`
  background: #232931;
  min-height: 100vh;
  font-family: "Montserrat", sans-serif;
`;
const PageTitle = styled.h1`
  font-size: 48px;
  color: #eeeeee;
  font-weight: 500;
`;
const PageContent = styled.div`
  padding: 2rem 3rem;
`;

const PageSubtitle = styled.h3`
  font-size: 24px;
  color: #eeeeee;
  font-weight: 400;
  margin-top: 3rem;
`;

const ResultContainer = styled(Row)`
  margin-top: 3rem;
`;

const ResultTextContainer = styled(Col)`
  background: #eeeeee;
  padding: 0.25rem 0.75rem;
`;

const ShortenerPage = () => {
  const [form] = Form.useForm();
  const [resultURL, setResultURL] = useState("");
  const clipboard = useClipboard();

  const shortenURL = async () => {
    const originalURL = form.getFieldValue("url");
    console.log(originalURL);

    // TODO: call api
    setResultURL("https://matcher.com/asdmgWa");
  };

  const copyResultURL = () => {
    clipboard.copy(resultURL);
    message.success({ content: "Copied to clipboard", duration: 1 });
  };

  useEffect(() => {
    if (resultURL) {
      copyResultURL();
    }
  }, [resultURL]);

  return (
    <PageContainer>
      <Navbar />
      <PageContent>
        <Form onFinish={shortenURL} form={form}>
          <PageTitle>Shorten your URL</PageTitle>
          <Row>
            <Col md={{ span: 8 }} xs={{ span: 24 }}>
              <Form.Item name="url">
                <StyledInput placeholder="Please input your url" />
              </Form.Item>
            </Col>
          </Row>
          <Button onClick={shortenURL}>Shorten my URL</Button>
        </Form>
        <PageSubtitle>Result: </PageSubtitle>
        <Row>
          <ResultTextContainer md={{ span: 8 }} xs={{ span: 24 }}>
            <Row justify="space-between">
              <Col>
                <a href={resultURL}>{resultURL}</a>
              </Col>
              <Col>
                <CopyOutlined onClick={copyResultURL} />
              </Col>
            </Row>
          </ResultTextContainer>
        </Row>
      </PageContent>
    </PageContainer>
  );
};

export default ShortenerPage;