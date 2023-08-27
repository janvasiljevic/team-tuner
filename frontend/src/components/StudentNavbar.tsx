import useLogout from '@/hooks/useLogout';
import { ActionIcon, Text, Flex, Container } from '@mantine/core';
import { IconLogout } from '@tabler/icons-react';
import React from 'react';

const StudentNavbar = () => {
  const { handleLogout } = useLogout();

  return (
    <Container w="100%">
      <Flex py="md" justify="space-between" align="center">
        <Text c="dimmed" size="lg">
          logged in as{' '}
          <Text span c="grape">
            janvasiljevic
          </Text>
        </Text>
        <ActionIcon size="lg" c="red" onClick={handleLogout}>
          <IconLogout />
        </ActionIcon>
      </Flex>
    </Container>
  );
};

export default StudentNavbar;
