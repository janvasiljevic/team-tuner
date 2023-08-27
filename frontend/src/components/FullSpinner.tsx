import { Flex, Loader } from '@mantine/core';
import React from 'react';

const FullSpinner = () => {
  return (
    <Flex h="100%" w="100%" align="center" justify="center">
      <Loader />
    </Flex>
  );
};

export default FullSpinner;
