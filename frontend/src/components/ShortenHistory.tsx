import { Row, Col, Timeline, message, Input } from 'antd'
import { useClipboard } from 'use-clipboard-copy'
import { CopyOutlined } from '@ant-design/icons'
import { getShortenHistory } from 'utils/api'
import styled from '@emotion/styled'
import { useEffect, useState } from 'react'

const { Search } = Input

const PageContainer = styled.div`
  background: #232931;
  font-family: 'Montserrat', sans-serif;
  padding: 0rem 3rem 2rem 3rem;
  min-height: 50vh;
  color: #fff;

  .ant-timeline {
    color: rgba(255, 255, 255, 0.65);
  }
  .ant-timeline-item-tail {
    border-left: 2px solid #303030;
  }
`

const TimelineItem = ({ color, timestamp, isp, ipType, region }) => {
  return (
    <Timeline.Item color={color}>
      {color == 'red' ? (
        'Response failed'
      ) : (
        <>
          <p>Timestamp: {timestamp}</p>
          <p>ISP: {isp}</p>
          <p>Type: {ipType}</p>
          <p>Region: {region}</p>
        </>
      )}
    </Timeline.Item>
  )
}

const History = () => {
  const clipboard = useClipboard()
  const resultURL = 'https://ant.design/components/timeline'

  const [history, setHistory] = useState(null)

  const copyResultURL = () => {
    clipboard.copy(resultURL)
    message.success({ content: 'Copied to clipboard', duration: 1 })
  }

  const onSearch = (val) => {
    console.log(val)
  }

  useEffect(() => {
    setHistory(getShortenHistory())
  }, [])

  return (
    <div style={{ marginBottom: '24px' }}>
      <Search
        placeholder="Search with shorten key"
        onSearch={onSearch}
        style={{ width: 400 }}
        size="large"
      />
      <div style={{ marginBottom: '24px' }}></div>

      {history ? (
        <>
          <Row>
            <Col>
              <h2 style={{ marginBottom: '16px', color: '#fff' }}>
                Shorten URL: {resultURL}
              </h2>
            </Col>
            <Col style={{ position: 'relative', top: '6px', right: '-12px' }}>
              <CopyOutlined onClick={copyResultURL} />
            </Col>
          </Row>
          <Timeline>
            <TimelineItem
              color="green"
              timestamp="2015-09-01 09:12:11"
              isp="AIS Fibre"
              ipType="IPv6"
              region="Bangkok"
            />
            <TimelineItem
              color="green"
              timestamp="2015-09-01 09:12:11"
              isp="AIS Fibre"
              ipType="IPv6"
              region="Bangkok"
            />
            <TimelineItem color="red" timestamp="" isp="" ipType="" region="" />
          </Timeline>
        </>
      ) : (
        <h2 style={{ marginBottom: '16px', color: 'rgb(130 130 130)' }}>
          No history found
        </h2>
      )}
    </div>
  )
}

export const ShortenHistory = () => {
  return (
    <PageContainer>
      <History />
    </PageContainer>
  )
}
