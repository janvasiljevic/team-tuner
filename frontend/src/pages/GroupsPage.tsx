import { useGetCoursesGroupRuns } from '@/api/courses/courses';
import {
  useDownloadGroupsCsv,
  useGetGroupById,
  useGetGroupsByGroupRun,
} from '@/api/groups/groups';
import { OutStudentOut } from '@/api/model';
import BfiBoxPlot from '@/components/BfiBoxPlot';
import { useAppStore } from '@/store';
import {
  Alert,
  Box,
  Button,
  Center,
  Chip,
  Container,
  Divider,
  Flex,
  Grid,
  Group,
  Loader,
  Paper,
  ScrollArea,
  Stack,
  Text,
  Title,
  UnstyledButton,
} from '@mantine/core';
import { IconAlertCircle, IconDownload } from '@tabler/icons-react';
import React, { useEffect } from 'react';

const GroupsPage = () => {
  const { selectedCourse } = useAppStore();

  const { data, isLoading } = useGetCoursesGroupRuns(selectedCourse?.id || '');
  const [selectedRunId, setSelectedRunId] = React.useState<string | null>(null);

  useEffect(() => {
    if (data && data.length > 0) {
      setSelectedRunId(data[0].id);
    }
  }, [data]);

  if (isLoading || !data) {
    return (
      <Center w="100%" h="100%">
        <Loader />
      </Center>
    );
  }

  if (data.length === 0) {
    return (
      <Center w="100%" h="100%">
        <Alert
          icon={<IconAlertCircle size="1rem" />}
          title="Bummer!"
          color="red"
        >
          No group formations found for course {selectedCourse?.code}
        </Alert>
      </Center>
    );
  }

  return (
    <Container w="100%" sx={{ fleexGrow: 1 }}>
      <Flex
        w="100%"
        h="100%"
        direction="column"
        sx={{ overflow: 'hidden' }}
        pb="lg"
      >
        <Paper mt="lg" p="lg">
          <Stack spacing="2px" mb="sm">
            <Title order={4} fw="900" c="blue">
              Group formations
            </Title>
            <Text size="sm" c="dimmed">
              {' '}
              for course {selectedCourse?.code}
            </Text>
          </Stack>
          <Flex wrap="wrap" gap="sm" justify="center">
            {data.map((groupRun) => (
              <Chip
                key={groupRun.id}
                variant="outline"
                size="xs"
                checked={groupRun.id === selectedRunId}
                onClick={() => setSelectedRunId(groupRun.id)}
              >
                {new Date(groupRun.createdAt).toLocaleString()}
              </Chip>
            ))}
          </Flex>
        </Paper>
        <GroupsSubpagePage groupRunId={selectedRunId} />
      </Flex>
    </Container>
  );
};

type SubPageProps = {
  groupRunId?: string | null;
};

const GroupsSubpagePage = ({ groupRunId }: SubPageProps) => {
  const [selectedGroupId, setSelectedGroupId] = React.useState<string | null>(
    null,
  );

  const { data: groupsData, isLoading: groupsLoading } = useGetGroupsByGroupRun(
    { groupRun: groupRunId || '' },
    { query: { enabled: !!groupRunId } },
  );

  const { refetch: downloadCsvFromServer } = useDownloadGroupsCsv(
    groupRunId || '',
    {
      query: { enabled: !!groupRunId },
    },
  );

  // Information about a single team (selected by the user)
  const { data: singleGroupData, isLoading: singleGroupLoading } =
    useGetGroupById(selectedGroupId || '', {
      query: { enabled: selectedGroupId !== null },
    });

  useEffect(() => {
    if (groupsData && groupsData.length > 0) {
      setSelectedGroupId(groupsData[0].id);
    }
  }, [groupsData]);

  const downloadCSV = async () => {
    const data = await downloadCsvFromServer();

    console.log(data.data);
  };

  return (
    <Grid sx={{ flexGrow: 1, overflow: 'hidden' }}>
      <Grid.Col span={6} h="100%">
        <Paper mt="lg" p="lg" h="100%">
          <Group position="right" w="100%">
            <Button
              size="xs"
              variant="light"
              leftIcon={<IconDownload />}
              color="green"
              onClick={downloadCSV}
            >
              Download CSV
            </Button>
          </Group>
          <ScrollArea mih="100%" h="100%">
            <Stack spacing="xs">
              {(groupsData || []).map((group, i) => (
                <UnstyledButton
                  key={group.id}
                  onClick={() => setSelectedGroupId(group.id)}
                  w="100%"
                >
                  <Flex w="100%" direction="column">
                    <Text ff="monospace">
                      {i + 1}. {group.name}
                    </Text>
                    <Text size="sm" c="dimmed">
                      {group.students.length} students
                    </Text>
                    <Divider w="100%" mt="xs" />
                  </Flex>
                </UnstyledButton>
              ))}
            </Stack>
          </ScrollArea>
        </Paper>
      </Grid.Col>
      <Grid.Col span={6} h="100%">
        <Paper mt="lg" p="lg" h="100%">
          <Stack spacing="xs">
            <Text fw="bold">{singleGroupData?.name || 'loading'}</Text>
            <Box h="300px" w="100%">
              <BfiBoxPlot
                showLabels={false}
                serverData={singleGroupData?.bigFiveBoxPlot}
              />
            </Box>
            <Text fw="bold" size="sm">
              Students
            </Text>
            <Stack>
              {singleGroupData?.students.map((student) => (
                <StudentStats key={student.id} student={student} />
              ))}
            </Stack>
          </Stack>
        </Paper>
      </Grid.Col>
    </Grid>
  );
};

const keys = [
  {
    key: 'Openness',
    color: 'blue',
  },
  {
    key: 'Concientiousness',
    color: 'red',
  },

  {
    key: 'Extraversion',
    color: 'yellow',
  },
  {
    key: 'Agreeableness',
    color: 'green',
  },
  {
    key: 'Neuroticism',
    color: 'purple',
  },
];

const StudentStats = ({ student }: { student: OutStudentOut }) => {
  return (
    <Stack w="100%" spacing="1px">
      <Text> {student.githubUsername} </Text>
      <Group w="100%" position="apart">
        {keys.map((key) => (
          <Stack key={key.key} spacing="2px">
            <Text size="sm" fw="bolder" c={key.color}>
              {(
                student[key.key.toLowerCase() as keyof OutStudentOut] as number
              ).toFixed(2)}
            </Text>
            <Text size="sm" c="dimmed">
              {key.key.charAt(0).toUpperCase()}
            </Text>
          </Stack>
        ))}
      </Group>
    </Stack>
  );
};

export default GroupsPage;
