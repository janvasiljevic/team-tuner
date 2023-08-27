import { AXIOS_INSTANCE } from '@/api/mutator/custom-instance';
import { showNotification } from '@mantine/notifications';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const useAxiosAuthInterceptorHook = () => {
  const navigate = useNavigate();
  const [isMounted, setIsMounted] = useState(false);

  // Mount the response interceptor
  useEffect(() => {
    const interceptor = AXIOS_INSTANCE.interceptors.response.use(
      (r) => r,
      (error) => {
        const status = error.response.status;

        if (status === 401) {
          showNotification({
            title: 'Unauthorized',
            message: 'You are not logged in',
          });

          navigate('/');
        }

        if (status === 403) {
          showNotification({
            title: 'Unauthorized',
            message: "You dont't have permission to access this page",
          });

          navigate('/');
        }

        return Promise.reject(error);
      },
    );

    setIsMounted(true);

    return () => {
      setIsMounted(false);
      AXIOS_INSTANCE.interceptors.response.eject(interceptor);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return isMounted;
};

export default useAxiosAuthInterceptorHook;
