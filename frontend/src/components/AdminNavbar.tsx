import useLogout from '@/hooks/useLogout';
import {
  ActionIcon,
  Box,
  Container,
  Group,
  Menu,
  SegmentedControl,
  Stack,
  Text,
  ThemeIcon,
  Title,
  UnstyledButton,
  createStyles,
} from '@mantine/core';
import { modals } from '@mantine/modals';
import { IconLogout, IconPlus, IconSelector } from '@tabler/icons-react';
import { useCallback, useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useGetCourse } from '../api/courses/courses';
import { CourseViewParams } from '../router';
import { useAppStore } from '../store';

const useStyles = createStyles((t) => ({
  header: {
    paddingTop: t.spacing.sm,
    backgroundColor: t.colors.gray[0],
  },
  mainSection: {
    paddingBottom: t.spacing.sm,
  },
}));

// This component wraps navbars. It used so that we can have a consistent
// heeader spacing and background color across all pages.
const AdminNavbar = () => {
  const navigate = useNavigate();
  const { selectedCourse, setSelectedCourse } = useAppStore();
  const { handleLogout } = useLogout();

  const { classes } = useStyles();

  const { data: coursesData } = useGetCourse({
    query: {
      queryKey: ['courses'],
    },
  });

  const { courseCode } = useParams<CourseViewParams>();

  const [segmentedControlValue, setSegmentedControlValue] =
    useState('students');

  const setSegmentedInterceptor = (value: string) => {
    setSegmentedControlValue(value);

    if (value === 'students') {
      navigate(`/admin/${selectedCourse?.code}/students`);
    } else {
      navigate(`/admin/${selectedCourse?.code}/groups`);
    }
  };

  const navigateToDefault = useCallback(() => {
    if (!coursesData) return;

    if (coursesData.length === 0) return;

    navigate(`/admin/${coursesData[0].code}/students`);
    setSelectedCourse(coursesData[0]);
  }, [coursesData, navigate, setSelectedCourse]);

  useEffect(() => {
    if (!coursesData) return;

    if (!courseCode) return navigateToDefault();

    const course = coursesData.find((c) => c.code === courseCode);

    if (!course) return navigateToDefault();

    setSelectedCourse(course);
  }, [coursesData, courseCode, navigateToDefault, setSelectedCourse]);

  return (
    <Box className={classes.header}>
      <Container size="xl" className={classes.mainSection}>
        <Group position="apart">
          <Menu>
            <Menu.Target>
              <UnstyledButton>
                <Group>
                  <Stack spacing="2px">
                    <Title
                      size="12px"
                      weight={700}
                      sx={{ color: selectedCourse?.colour }}
                    >
                      {selectedCourse?.code}
                    </Title>
                    <Text c="dimmed" size="sm" ff="monospace" weight={700}>
                      {selectedCourse?.name}
                    </Text>
                  </Stack>
                  <IconSelector />
                </Group>
              </UnstyledButton>
            </Menu.Target>
            <Menu.Dropdown>
              {coursesData?.map((course) => (
                <Menu.Item
                  key={course.id}
                  onClick={() => {
                    setSelectedCourse(course);
                    navigate(`/admin/${course.code}/students`);
                  }}
                >
                  <Group>
                    <Stack spacing="2px">
                      <Title
                        size="12px"
                        weight={700}
                        sx={{ color: course.colour }}
                      >
                        {course.code}
                      </Title>
                      <Text c="dimmed" size="sm" ff="monospace" weight={700}>
                        {course.name}
                      </Text>
                    </Stack>
                  </Group>
                </Menu.Item>
              ))}
              <Menu.Divider />
              <Menu.Item
                onClick={() => {
                  modals.openContextModal({
                    modal: 'addCourse',
                    title: 'Add a new course',
                    innerProps: {
                      nav: navigate,
                    },
                  });
                }}
                icon={
                  <ThemeIcon variant="light" color="blue">
                    <IconPlus />
                  </ThemeIcon>
                }
              >
                Add a new course
              </Menu.Item>
            </Menu.Dropdown>
          </Menu>
          <Group spacing="8rem">
            <SegmentedControl
              color="blue"
              onChange={setSegmentedInterceptor}
              value={segmentedControlValue}
              data={[
                { label: 'Students', value: 'students' },
                { label: 'Groups', value: 'groups' },
              ]}
            />

            <ActionIcon color="red" onClick={handleLogout}>
              <IconLogout />
            </ActionIcon>
          </Group>
        </Group>
      </Container>
    </Box>
  );
};

export default AdminNavbar;
