import {
  Blockquote,
  Box,
  Button,
  Container,
  Flex,
  Group,
  Rating,
  Skeleton,
  Stack,
  Text,
  Title,
} from '@mantine/core';
import { modals } from '@mantine/modals';
import { useState } from 'react';
import { GetQuestionTypeOfQuestion, OutQuestioneItemOut } from '../api/model';
import {
  useGetQuestion,
  usePostQuestionAnswer,
  usePostQuestionSubmit,
} from '../api/questions/questions';

const steps: GetQuestionTypeOfQuestion[] = [
  'agreeableness',
  'conscientiousness',
  'extroversion',
  'neuroticism',
  'openness',
];

const QuestionPage = () => {
  const [stepIndex, setStepIndex] = useState(0);

  const { data, isLoading, refetch } = useGetQuestion({
    typeOfQuestion: steps[stepIndex],
  });

  const { mutateAsync } = usePostQuestionAnswer();
  const { mutateAsync: submitAnswers } = usePostQuestionSubmit();

  const answer = async (answerId: string, answerValue: number) => {
    await mutateAsync({
      data: { answerID: answerId, value: answerValue },
    });

    refetch();
  };

  const openModal = (unresolved_questions: OutQuestioneItemOut[]) =>
    modals.open({
      title: 'Missing questions',
      centered: true,
      children: (
        <Box my="sm">
          <Text fw="700">
            You have not answered all questions. Missing{' '}
            {unresolved_questions.length} answers
          </Text>
          <Stack mt="lg">
            {unresolved_questions.slice(0, 5).map((question) => (
              <li key={question.question_id}> {question.question}</li>
            ))}
            {unresolved_questions.length > 5 && <li> ... </li>}
          </Stack>
        </Box>
      ),
    });

  const nextStep = async () => {
    if (stepIndex < steps.length - 1) {
      setStepIndex(stepIndex + 1);
    } else if (stepIndex === steps.length - 1) {
      const d = await submitAnswers();

      if (d.unresolved_questions.length > 0) {
        openModal(d.unresolved_questions);
      }
    }
  };

  const prevStep = () => {
    if (stepIndex > 0) {
      setStepIndex(stepIndex - 1);
    }
  };

  return (
    <Container w="100%">
      <Box h="8rem" />
      {isLoading && (
        <>
          <Skeleton>
            <Title>loading</Title>
          </Skeleton>
          <Skeleton mb="4rem">
            <Blockquote>loading</Blockquote>
          </Skeleton>
          <Stack>
            {Array.from({ length: 6 }).map((_, i) => (
              <Skeleton key={i} mb="1rem">
                <Text>loading</Text>
              </Skeleton>
            ))}
          </Stack>
        </>
      )}

      {data && !isLoading && (
        <>
          <Title>{data.title}</Title>
          <Blockquote mb="4rem"> {data.description}</Blockquote>
          <Stack>
            {data.questions.map((question) => (
              <QuestionItem
                key={question.question_id}
                question={question}
                answer={answer}
              />
            ))}
          </Stack>
        </>
      )}

      <Group mt="5rem" position="right" w="100%">
        <Button variant="light" onClick={prevStep} disabled={stepIndex === 0}>
          Previous
        </Button>
        <Button variant="outline" onClick={nextStep}>
          {stepIndex === steps.length - 1 ? 'Submit' : 'Next'}
        </Button>
      </Group>
    </Container>
  );
};

type QuestionItemProps = {
  question: OutQuestioneItemOut;
  answer: (questionId: string, answerValue: number) => void;
};

const QuestionItem = ({ question, answer }: QuestionItemProps) => {
  const [value, setValue] = useState(question.answer_value || 0);

  const onChange = (value: number) => {
    setValue(value);
    answer(question.answer_id, value);
  };

  return (
    <Flex justify="space-between">
      <Text> {question.question} </Text>
      <Rating color="green" value={value} onChange={onChange} />
    </Flex>
  );
};

export default QuestionPage;
