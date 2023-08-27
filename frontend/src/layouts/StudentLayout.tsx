import StudentNavbar from '@/components/StudentNavbar';
import { Flex } from '@mantine/core';
import React from 'react';
import { useOutlet } from 'react-router-dom';

const StudentLayout = () => {
  const outlet = useOutlet();

  return (
    <Flex w="100%" h="100vh" direction="column" sx={{ overflow: 'hidden' }}>
      <StudentNavbar />
      <Flex
        sx={{
          flexGrow: 1,
          overflowX: 'hidden',
          overflowY: 'auto',
        }}
        h="100%"
        w="100%"
      >
        {outlet}
      </Flex>
    </Flex>
  );
};

export default StudentLayout;
