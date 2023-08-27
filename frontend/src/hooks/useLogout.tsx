import { usePostAuthLogout } from '@/api/auth/auth';
import { showNotification } from '@mantine/notifications';
import { IconLogout } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

const useLogout = () => {
  const { mutateAsync: logout } = usePostAuthLogout();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();

    showNotification({
      title: 'Logged out',
      message: 'You have been logged out',
      icon: <IconLogout />,
    });

    navigate('/');
  };

  return { handleLogout };
};

export default useLogout;
