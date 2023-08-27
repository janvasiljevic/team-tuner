import {
  useGetCourseCourseIdStatsBfBoxPlot,
  useGetCourseCourseIdStatsQuestionerStats,
} from '@/api/courses/courses';
import {
  Button,
  Card,
  Container,
  Flex,
  Grid,
  Group,
  Stack,
  Text,
} from '@mantine/core';
import { IconActivity, IconPentagon, IconUsers } from '@tabler/icons-react';
import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
  type MRT_PaginationState,
} from 'mantine-react-table';
import { useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { OutStudentOut } from '../api/model';
import { useGetStudent } from '../api/students/students';
import BfiBoxPlot from '../components/BfiBoxPlot';
import DailyActivity from '../components/DailyActivity';
import { useAppStore } from '../store';

const AdminStudentsPage = () => {
  const { selectedCourse } = useAppStore();

  const { data: serverData } = useGetCourseCourseIdStatsBfBoxPlot(
    selectedCourse?.id || '',
    {
      query: {
        enabled: !!selectedCourse?.id,
      },
    },
  );

  return (
    <Container size="xl" w="100%" h="100%">
      <Flex direction="column">
        <Grid h="300px" mt="lg">
          <GraphWrapper>
            <Group spacing="xs" w="100%" position="center">
              <IconPentagon />
              <Text>BF Overview</Text>
            </Group>
            <BfiBoxPlot serverData={serverData} />
          </GraphWrapper>
          <GraphWrapper>
            <Group spacing="xs" w="100%" position="center">
              <IconActivity />
              <Text>Daily activity</Text>
            </Group>
            <DailyActivity />
          </GraphWrapper>
        </Grid>
        <GroupsToolbar />
        <Card mt="lg">
          <StudentsTable />
        </Card>
      </Flex>
    </Container>
  );
};

const GroupsToolbar = () => {
  const { selectedCourse } = useAppStore();

  const navigate = useNavigate();

  const { data } = useGetCourseCourseIdStatsQuestionerStats(
    selectedCourse?.id || '',
    {
      query: {
        enabled: !!selectedCourse,
      },
    },
  );

  const handleFormGroups = async () => {
    if (!selectedCourse) return;

    navigate(`/admin/${selectedCourse.code}/groups/new`);
  };

  return (
    <Card mt="lg">
      <Flex justify="space-between" align="center">
        <Group>
          <IconUsers size={24} />
          <Stack spacing="2px">
            <Text fw="700">
              {' '}
              {data?.completed || 0} / {data?.totalCount || 0}
            </Text>
            <Text size="xs"> questioners solved</Text>
          </Stack>
        </Group>
        <Button onClick={handleFormGroups}> Form groups</Button>
      </Flex>
    </Card>
  );
};

const GraphWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <Grid.Col span={6} h="300px">
      <Card w="100%" h="100%">
        <Flex direction="column" h="100%" w="100%">
          {children}
        </Flex>
      </Card>
    </Grid.Col>
  );
};

const StudentsTable = () => {
  const { selectedCourse } = useAppStore();

  const [pagination, setPagination] = useState<MRT_PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const columns = useMemo<MRT_ColumnDef<OutStudentOut>[]>(
    () => [
      {
        accessorKey: 'githubUsername',
        header: 'GH Username',
        enableColumnFilter: false,
      },
      {
        accessorKey: 'agreeableness',
        header: 'Agreeableness',
        enableColumnFilter: false,
        accessorFn: (row) => {
          if (row.agreeableness) {
            return row.agreeableness.toFixed(2);
          }
        },
      },
      {
        accessorKey: 'neuroticism',
        header: 'Neuroticism',
        enableColumnFilter: false,
        accessorFn: (row) => {
          if (row.neuroticism) {
            return row.neuroticism.toFixed(2);
          }
        },
      },
      {
        accessorKey: 'extraversion',
        header: 'Extraversion',
        enableColumnFilter: false,
        accessorFn: (row) => {
          if (row.extraversion) {
            return row.extraversion.toFixed(2);
          }
        },
      },
      {
        accessorKey: 'concientiousness',
        header: 'Conscientiousness',
        enableColumnFilter: false,
        accessorFn: (row) => {
          if (row.concientiousness) {
            return row.concientiousness.toFixed(2);
          }
        },
      },
      {
        accessorKey: 'openness',
        header: 'Openness',
        // set float to 2 decimal places
        enableColumnFilter: false,
        accessorFn: (row) => {
          if (row.openness) {
            return row.openness.toFixed(2);
          }
        },
      },
    ],
    [],
  );

  const { data, isLoading } = useGetStudent(
    {
      courseId: selectedCourse?.id || '',
      page: pagination.pageIndex,
      pageSize: pagination.pageSize,
    },
    {
      query: {
        enabled: !!selectedCourse,
        keepPreviousData: true,
      },
    },
  );

  const table = useMantineReactTable({
    columns,
    paginationDisplayMode: 'pages',
    data: data?.content || [],
    enableTopToolbar: false,
    manualPagination: true,
    enableColumnFilterModes: false,
    rowCount: data?.totalCount || 0,
    mantinePaperProps: {
      shadow: 'none',
    },
    state: {
      pagination,
      isLoading,
      density: 'xs',
      columnPinning: {
        left: ['githubUsername'],
      },
    },
    enableGlobalFilter: false,
    onPaginationChange: setPagination,
    enablePinning: true,
  });

  return (
    <>
      <MantineReactTable table={table} />
    </>
  );
};

export default AdminStudentsPage;
