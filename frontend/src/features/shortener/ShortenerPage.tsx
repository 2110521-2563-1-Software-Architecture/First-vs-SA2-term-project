import styled from '@emotion/styled'
import { Button, Form, Input, Row, Col } from 'antd'
import Navbar from 'components/Navbar'

const StyledInput = styled(Input)`
  border-radius: 0;
  padding: 0.5rem 1rem;
  font-family: 'Montserrat', sans-serif;
`
const PageContainer = styled.div`
  background: #232931;
  min-height: 100vh;
  font-family: 'Montserrat', sans-serif;
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
`

const ResultContainer = styled(Row)`
  margin-top: 3rem;
`

const ResultTextContainer = styled(Col)`
  background: #eeeeee;
  padding: 0.5rem 1rem;
`

const ShortenerPage = () => {
  const [form] = Form.useForm()

  const shortenURL = () => {
    const originalURL = form.getFieldValue('url')
    console.log(originalURL)

    // TODO: call api
  }

  return (
    <PageContainer>
      <Navbar />
      <PageContent>
        <Form onFinish={shortenURL} form={form}>
          <PageTitle>Shorten your URL</PageTitle>
          <Row>
            <Col span={8}>
              <Form.Item name="url">
                <StyledInput placeholder="Please input your url" />
              </Form.Item>
            </Col>
          </Row>
          <Button onClick={shortenURL}>Shorten my URL</Button>
        </Form>
        <ResultContainer align="middle">
          <Col span={2}>
            <PageSubtitle>Result: </PageSubtitle>
          </Col>
          <ResultTextContainer span={8}>result</ResultTextContainer>
        </ResultContainer>
      </PageContent>
    </PageContainer>
  )
}

export default ShortenerPage
