import { useGetAuthWhoami } from '@/api/auth/auth';
import FullSpinner from '@/components/FullSpinner';
import { useEffect } from 'react';
import { useNavigate } from 'react-router';

const LoggedInPage = () => {
  const { data } = useGetAuthWhoami();
  const navigate = useNavigate();

  useEffect(() => {
    if (!data) return;

    if (data.role === 'student') {
      if (data.finished_bfi) navigate('/report');
      else navigate('/questions');
    } else if (data.role === 'admin') {
      navigate('/admin');
    }
  }, [data, navigate]);

  return <FullSpinner />;
};

export default LoggedInPage;
