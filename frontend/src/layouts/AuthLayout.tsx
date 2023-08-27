import FullSpinner from '@/components/FullSpinner';
import useAxiosAuthInterceptorHook from '@/hooks/useAxiosAuthInterceptor';
import { useOutlet } from 'react-router-dom';

const AuthLayout = () => {
  const isMounted = useAxiosAuthInterceptorHook();
  const outlet = useOutlet();

  if (!isMounted) return <FullSpinner />;

  return <>{outlet}</>;
};

export default AuthLayout;
