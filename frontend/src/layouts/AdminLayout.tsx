import { Flex } from '@mantine/core';
import { useOutlet } from 'react-router-dom';
import AdminNavbar from '../components/AdminNavbar';

const AdminLayout = () => {
  const outlet = useOutlet();

  return (
    <Flex
      w="100%"
      h="100vh"
      direction="column"
      sx={(s) => ({ overflow: 'hidden', background: s.colors.gray[2] })}
    >
      <AdminNavbar />
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

export default AdminLayout;
