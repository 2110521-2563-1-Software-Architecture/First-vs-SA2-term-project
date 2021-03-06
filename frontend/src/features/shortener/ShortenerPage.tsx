import styled from '@emotion/styled'
import { Button, Form, Input, Row, Col, message } from 'antd'
import { CopyOutlined } from '@ant-design/icons'
import { useClipboard } from 'use-clipboard-copy'
import { useEffect, useState } from 'react'
import Lottie from 'react-lottie'
import animationData from './animation.json'
import { getShortenURL, redirectTo } from 'utils/api'
import { NEXT_PUBLIC_SERVER } from 'utils/config'

const StyledInput = styled(Input)`
  border-radius: 0;
  padding: 0.25rem 0.75rem;
  font-family: 'Montserrat', sans-serif;
`
const PageContainer = styled.div`
  background: #232931;
  font-family: 'Montserrat', sans-serif;
  padding-top: 3rem;
`
const PageTitle = styled.h1`
  font-size: 48px;
  color: #eeeeee;
  font-weight: 500;
`
const PageContent = styled.div`
  padding: 2rem 3rem;
`

const PageSubtitle = styled.h3`
  font-size: 24px;
  color: #eeeeee;
  font-weight: 400;
  margin-top: 3rem;
`

const ResultContainer = styled(Row)`
  margin-top: 3rem;
`

const ResultTextContainer = styled(Col)`
  background: #eeeeee;
  padding: 0.25rem 0.75rem;
`

const ShortenerPage = () => {
  const [form] = Form.useForm()
  const [resultURL, setResultURL] = useState('')
  const clipboard = useClipboard()

  const defaultOptions = {
    loop: true,
    autoplay: true,
    animationData: animationData,
  }

  const shortenURL = async () => {
    const originalURL = form.getFieldValue('url')
    console.log(originalURL)

    // TODO: call api
    const { key } = await getShortenURL(originalURL)
    setResultURL(`${NEXT_PUBLIC_SERVER}/${key}`)
  }

  const copyResultURL = () => {
    clipboard.copy(resultURL)
    message.success({ content: 'Copied to clipboard', duration: 1 })
  }

  const redirect = async () => {
    const res = await redirectTo(resultURL)
  }

  useEffect(() => {
    if (resultURL) {
      copyResultURL()
    }
  }, [resultURL])

  return (
    <PageContainer>
      <Row>
        <Col md={{ span: 12 }} xs={{ span: 24 }}>
          <PageContent>
            <Form onFinish={shortenURL} form={form}>
              <PageTitle>Shorten your URL</PageTitle>
              <Row>
                <Col md={{ span: 16 }} xs={{ span: 24 }}>
                  <Form.Item name="url">
                    <StyledInput
                      size="large"
                      placeholder="Please input your url"
                    />
                  </Form.Item>
                </Col>
              </Row>
              <Button onClick={shortenURL}>Shorten my URL</Button>
            </Form>
            <PageSubtitle>Result: </PageSubtitle>
            <Row>
              <ResultTextContainer md={{ span: 16 }} xs={{ span: 24 }}>
                <Row justify="space-between">
                  <Col>
                    <span onClick={redirect}>{`${resultURL}`}</span>
                  </Col>
                  <Col>
                    <CopyOutlined onClick={copyResultURL} />
                  </Col>
                </Row>
              </ResultTextContainer>
            </Row>
          </PageContent>
        </Col>
        <Col md={{ span: 12 }} xs={{ span: 0 }}>
          <Lottie
            options={defaultOptions}
            height={400}
            width={400}
            style={{ marginTop: '2rem' }}
          />
        </Col>
      </Row>

      <PageContent>
        <br />
        <br />
        <PageTitle>Shorten Traffic History</PageTitle>
      </PageContent>
    </PageContainer>
  )
}

export default ShortenerPage
