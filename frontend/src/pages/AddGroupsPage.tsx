import { useGetCourseCourseIdStatsQuestionerStats } from '@/api/courses/courses';
import { useCreateGroups } from '@/api/groups/groups';
import { useAppStore } from '@/store';
import {
  Accordion,
  ActionIcon,
  Box,
  Button,
  Card,
  Center,
  Collapse,
  Container,
  Divider,
  Flex,
  Group,
  Loader,
  NumberInput,
  NumberInputHandlers,
  Stack,
  Text,
  ThemeIcon,
  UnstyledButton,
  createStyles,
  rem,
  useMantineTheme,
} from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import {
  IconAlertCircle,
  IconCheck,
  IconMinus,
  IconPlus,
  IconSettings2,
  IconTrash,
  IconUpload,
  IconUsersGroup,
} from '@tabler/icons-react';
import { useMemo, useRef, useState } from 'react';

const useStyles = createStyles((t) => ({
  group: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    width: '40px',
    height: '40px',
    background: t.colors.gray[2],
    borderRadius: t.radius.md,
  },
}));

type AdvancedSettingsForm = {
  temperature: number;
  iterations: number;
  weightSatisfaction: number;
  weightNeuroticism: number;
  weightExtraversion: number;
  weightConscientiousness: number;
};

const AddGroupsPage = () => {
  const { selectedCourse } = useAppStore();
  const { mutateAsync: formServerGroups } = useCreateGroups({
    mutation: {
      onError: (error) => {
        showNotification({
          title: 'Error',
          message: error.message,
          color: 'red',
        });
      },
    },
    request: {
      paramsSerializer: {
        indexes: null,
      },
    },
  });

  const { data, isLoading } = useGetCourseCourseIdStatsQuestionerStats(
    selectedCourse?.id || '',
    {
      query: {
        enabled: !!selectedCourse,
      },
    },
  );

  const [groupSize, setGroupSize] = useState<number | ''>(5);
  const handlers = useRef<NumberInputHandlers>();

  const { classes: c } = useStyles();

  const theme = useMantineTheme();
  const getColor = (color: string) =>
    theme.colors[color][theme.colorScheme === 'dark' ? 5 : 7];

  const advancedSettingForm = useForm<AdvancedSettingsForm>({
    initialValues: {
      temperature: 12,
      iterations: 10000,
      weightConscientiousness: 1.5,
      weightExtraversion: 1,
      weightNeuroticism: 1,
      weightSatisfaction: 0.8,
    },
  });

  const [groups, setGroups] = useState<number[]>([]);
  const [selectedGroupIndex, setSelectedGroupIndex] = useState<number>(-1);

  // How many students aren't sorted in groups
  const remainingStudents = useMemo(() => {
    if (!data) return 0;

    let totalStudents = data.totalCount;

    for (let i = 0; i < groups.length; i++) {
      totalStudents -= groups[i];
    }

    return totalStudents;
  }, [data, groups]);

  // Distribute students evenly into groups
  const distributeStudents = () => {
    if (!data) return;

    if (groupSize === '') {
      return;
    }

    const length = Math.floor(data.completed / groupSize);

    setGroups(Array.from({ length }, () => groupSize));

    setSelectedGroupIndex(0);

    // from section 5.6 "Iterations tuning" from thesis
    if (length < 10) {
      advancedSettingForm.setFieldValue('iterations', 6000);
    } else {
      advancedSettingForm.setFieldValue('iterations', 710 * length - 1700);
    }
  };

  const addRemainingStudents = () => {
    if (!data) return;

    if (remainingStudents === 0) {
      return;
    }

    setGroups([...groups, remainingStudents]);
  };

  const deleteGroup = (index: number) => {
    if (index < 0 || index >= groups.length) return;

    const newGroups = [...groups];
    newGroups.splice(index, 1);

    setGroups(newGroups);
  };

  const incramentGroupCount = (index: number, incrament: number) => {
    if (index < 0 || index >= groups.length) return;

    const newGroups = [...groups];
    newGroups[index] += incrament;

    setGroups(newGroups);
  };

  const createGroups = async () => {
    await formServerGroups({
      data: {
        courseId: selectedCourse?.id || '',
        groupSizes: groups,
        ...advancedSettingForm.values,
      },
    });

    showNotification({
      title: 'Groups created',
      message: 'Groups have been created successfully',
      color: 'green',
    });
  };

  if (isLoading || !data) {
    return (
      <Center w="100%" h="100%">
        <Loader />
      </Center>
    );
  }

  return (
    <Container w="100%" size="md" pt="lg">
      <Card>
        <Stack w="100%">
          <Group>
            <ThemeIcon variant="filled">
              <IconCheck />
            </ThemeIcon>
            <Text>
              {data.completed} / {data.totalCount} students have completed the
              questioner
            </Text>
            <Text c="dimmed">
              After all the students have finished the questioner you can
              proceed. Select the desired group size and distribute the students
              evenly into groups. Manually assign students to groups by clicking
              on the group and tweaking the group size.
            </Text>
          </Group>
          <Group w="100%" position="apart">
            <Group spacing="xs">
              <Text mr="lg"> Group size</Text>
              <ActionIcon
                size="lg"
                variant="default"
                onClick={() => handlers?.current?.decrement()}
              >
                –
              </ActionIcon>

              <NumberInput
                hideControls
                value={groupSize}
                onChange={(val) => setGroupSize(val)}
                handlersRef={handlers}
                max={7}
                min={3}
                step={1}
                styles={{ input: { width: rem(54), textAlign: 'center' } }}
              />

              <ActionIcon
                size="lg"
                variant="default"
                onClick={() => handlers?.current?.increment()}
              >
                +
              </ActionIcon>
            </Group>

            <Button leftIcon={<IconUsersGroup />} onClick={distributeStudents}>
              Distribute students
            </Button>
          </Group>
          {groups.length !== 0 && (
            <Stack>
              <Group position="apart">
                <Text fw="700"> Distribution </Text>
              </Group>
              <Flex gap="sm" wrap="wrap" align="center">
                {groups.map((group, index) => (
                  <UnstyledButton
                    key={index}
                    onClick={() => setSelectedGroupIndex(index)}
                  >
                    <Box
                      className={c.group}
                      sx={(t) => ({
                        ...(selectedGroupIndex === index && {
                          border: `2px solid ${t.colors.blue[6]}`,
                        }),
                      })}
                    >
                      <Text> {group} </Text>
                    </Box>
                  </UnstyledButton>
                ))}

                <ActionIcon
                  size="lg"
                  disabled={remainingStudents === 0}
                  onClick={addRemainingStudents}
                  variant="subtle"
                >
                  <IconPlus size="16px" />
                </ActionIcon>
              </Flex>

              <Group position="apart">
                <Stack spacing="2px">
                  <Text>
                    Selected group
                    <Text span fw="900" c="blue">
                      {' '}
                      {selectedGroupIndex + 1}
                    </Text>
                  </Text>
                  <Text size="sm" c="dimmed">
                    with {groups[selectedGroupIndex]} students
                  </Text>
                </Stack>
                <Group>
                  <ActionIcon
                    c="blue"
                    variant="light"
                    size="lg"
                    disabled={groups[selectedGroupIndex] <= 1}
                    onClick={() => incramentGroupCount(selectedGroupIndex, -1)}
                  >
                    <IconMinus size="16px" stroke={4} />
                  </ActionIcon>

                  <ActionIcon
                    c="blue"
                    variant="light"
                    disabled={remainingStudents === 0}
                    size="lg"
                    onClick={() => incramentGroupCount(selectedGroupIndex, 1)}
                  >
                    <IconPlus size="16px" stroke={4} />
                  </ActionIcon>

                  <Box w={8} />

                  <ActionIcon
                    onClick={() => deleteGroup(selectedGroupIndex)}
                    c="red"
                    variant="light"
                    size="lg"
                  >
                    <IconTrash size="16px" />
                  </ActionIcon>
                </Group>
              </Group>

              <Divider />

              <Group position="apart">
                <Group>
                  {remainingStudents === 0 ? (
                    <ThemeIcon variant="light" color="green">
                      <IconCheck size="16px" />
                    </ThemeIcon>
                  ) : (
                    <ThemeIcon variant="light" color="red">
                      <IconAlertCircle size="16px" />
                    </ThemeIcon>
                  )}
                  <Text size="sm">
                    {remainingStudents} students not assigned to a group
                  </Text>
                </Group>
                <Button
                  disabled={remainingStudents !== 0}
                  color="green"
                  onClick={createGroups}
                  leftIcon={<IconUpload />}
                >
                  Generate groups
                </Button>
              </Group>

              <Accordion
                chevronPosition="left"
                w={'100%'}
                mx="auto"
                variant="contained"
              >
                <Accordion.Item value="item-advanced">
                  <Accordion.Control
                    icon={<IconSettings2 color={getColor('orange')} />}
                  >
                    Advanced settings
                  </Accordion.Control>
                  <Accordion.Panel>
                    <Text mb="sm">
                      Settings related to tuning the simulated annealing
                      algorithm.
                    </Text>
                    <form>
                      <Stack>
                        <Text size="xs" fw={500}>
                          General
                        </Text>
                        <Group>
                          <NumberInput
                            label="Iterations"
                            min={100}
                            max={30000}
                            step={100}
                            {...advancedSettingForm.getInputProps('iterations')}
                          />
                          <NumberInput
                            label="Temperature"
                            min={1}
                            max={100}
                            precision={2}
                            step={0.1}
                            {...advancedSettingForm.getInputProps(
                              'temperature',
                            )}
                          />
                        </Group>
                        <Text size="xs" fw={500}>
                          Weights
                        </Text>
                        <Group>
                          <NumberInput
                            label="Satisfaction"
                            min={0.3}
                            max={2}
                            precision={2}
                            step={0.05}
                            {...advancedSettingForm.getInputProps(
                              'weightSatisfaction',
                            )}
                          />
                          <NumberInput
                            label="Neuroticism"
                            min={0.3}
                            max={2}
                            step={0.05}
                            precision={2}
                            {...advancedSettingForm.getInputProps(
                              'weightNeuroticism',
                            )}
                          />
                          <NumberInput
                            label="Extroversion"
                            min={0.3}
                            max={2}
                            step={0.05}
                            precision={2}
                            {...advancedSettingForm.getInputProps(
                              'weightExtraversion',
                            )}
                          />
                          <NumberInput
                            label="Conscientiousness"
                            min={0.3}
                            max={2}
                            step={0.05}
                            precision={2}
                            {...advancedSettingForm.getInputProps(
                              'weightConscientiousness',
                            )}
                          />
                        </Group>
                      </Stack>
                    </form>
                  </Accordion.Panel>
                </Accordion.Item>
              </Accordion>
            </Stack>
          )}
        </Stack>
      </Card>
    </Container>
  );
};

export default AddGroupsPage;
