import { usePostCourse } from '@/api/courses/courses';
import { Button, ColorInput, Group, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { ContextModalProps } from '@mantine/modals';
import { showNotification } from '@mantine/notifications';
import { useQueryClient } from '@tanstack/react-query';
import { NavigateFunction } from 'react-router-dom';

type AddCourseModalForm = {
  code: string;
  name: string;
  colour: string;
};

const AddCourseModal = ({
  id,
  context,
  innerProps: { nav },
}: ContextModalProps<{ nav: NavigateFunction }>) => {
  const form = useForm<AddCourseModalForm>({
    initialValues: {
      code: '',
      name: '',
      colour: '#25262b',
    },
  });

  const qc = useQueryClient();

  const { mutateAsync: postData } = usePostCourse();

  const handleSubmit = async (values: AddCourseModalForm) => {
    await postData({ data: values })
      .then(async (res) => {
        showNotification({
          title: 'Success',
          message: 'Course added',
        });

        await qc.invalidateQueries(['courses']);

        context.closeContextModal(id);
        nav(`/admin/${res.code}/students`);
      })
      .catch((err) => {
        showNotification({
          title: 'Error',
          message: err.message,
        });
      });
  };

  return (
    <Group w="100%">
      <form onSubmit={form.onSubmit(handleSubmit)} style={{ width: '100%' }}>
        <Stack spacing="xs" mb="xl">
          <TextInput w="100%" label="Code" {...form.getInputProps('code')} />
          <TextInput w="100%" label="Name" {...form.getInputProps('name')} />
          <ColorInput
            placeholder="Pick color"
            label="Colour"
            w="100%"
            {...form.getInputProps('colour')}
          />
        </Stack>

        <Group position="right" w="100%">
          <Button
            variant="light"
            onClick={() => {
              context.closeContextModal(id);
            }}
          >
            Cancel
          </Button>
          <Button type="submit">Add</Button>
        </Group>
      </form>
    </Group>
  );
};

export default AddCourseModal;
