import { useGetAnswer, useGetAnswerBfiReport } from '@/api/answer/answer';
import {
  Box,
  Center,
  Container,
  Loader,
  Rating,
  Table,
  Tabs,
} from '@mantine/core';
import { ResponsiveRadar } from '@nivo/radar';
import { AnimatePresence, motion } from 'framer-motion';
import { useEffect, useMemo, useState } from 'react';
import { OutQuesitonOutType, OutQuestioneItemOut } from '../api/model';

const StudentReportPage = () => {
  return (
    <Container w="100%">
      <Box w="100%" h="500px">
        <Radar />
      </Box>
      <TableAnswers />
    </Container>
  );
};

const TableAnswers = () => {
  const { data } = useGetAnswer({
    query: {
      keepPreviousData: true,
    },
  });

  const [filteresData, setFilteredData] = useState<OutQuestioneItemOut[]>([]);

  const [activeTab, setActiveTab] = useState<OutQuesitonOutType | null>(
    'agreeableness',
  );

  useEffect(() => {
    if (!data) return;

    setFilteredData(data.filter((item) => item.question_type === activeTab));
  }, [data, activeTab]);

  return (
    <>
      <Tabs
        value={activeTab}
        onTabChange={(value) => {
          setActiveTab(value as OutQuesitonOutType);
        }}
        mb="lg"
      >
        <Tabs.List position="center">
          <Tabs.Tab value="agreeableness">Agreeablenes </Tabs.Tab>
          <Tabs.Tab value="conscientious">Conscientious </Tabs.Tab>
          <Tabs.Tab value="extraversion">Extraversion</Tabs.Tab>
          <Tabs.Tab value="neuroticism">Neuroticism</Tabs.Tab>
          <Tabs.Tab value="openness">Openness</Tabs.Tab>
        </Tabs.List>
      </Tabs>
      <Table>
        <thead>
          <tr>
            <th>Question</th>
            <th>Answer</th>
          </tr>
        </thead>
        <tbody>
          {filteresData.map((item) => (
            <tr key={item.question_id}>
              <td width="70%">{item.question}</td>
              <td width="30%">
                <Rating readOnly value={item.answer_value} />
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </>
  );
};

const Radar = () => {
  const { data: reportData, isLoading } = useGetAnswerBfiReport({});

  const reshapedData = useMemo(() => {
    if (!reportData) return [];

    return Object.entries(reportData).map(([facet, value]) => {
      return {
        facet,
        value,
      };
    });
  }, [reportData]);

  if (isLoading)
    return (
      <Center w="100%" h="100%">
        <Loader />
      </Center>
    );

  return (
    <AnimatePresence>
      {!isLoading && (
        <motion.div
          initial={{ opacity: 0, scale: 0.5 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.8 }}
          style={{ width: '100%', height: '100%' }}
        >
          <ResponsiveRadar
            data={reshapedData}
            keys={['value']}
            indexBy="facet"
            valueFormat=">-.2f"
            margin={{ top: 70, right: 40, bottom: 80, left: 40 }}
            gridLabelOffset={36}
            dotSize={10}
            dotColor={{ theme: 'background' }}
            dotBorderWidth={2}
            colors={{ scheme: 'category10' }}
            blendMode="multiply"
          />
        </motion.div>
      )}
    </AnimatePresence>
  );
};

export default StudentReportPage;
