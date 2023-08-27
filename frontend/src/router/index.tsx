import FullSpinner from '@/components/FullSpinner';
import AuthLayout from '@/layouts/AuthLayout';
import AddGroupsPage from '@/pages/AddGroupsPage';
import GroupsPage from '@/pages/GroupsPage';
import StudentReportPage from '@/pages/StudentReportPage';
import { createBrowserRouter } from 'react-router-dom';
import AdminLayout from '../layouts/AdminLayout';
import AdminStudentsPage from '../pages/AdminStudentsPage';
import LoggedInPage from '../pages/LoggedInPage';
import LoginPage from '../pages/LoginPage';
import QuestionPage from '../pages/QuestionPage';
import StudentLayout from '@/layouts/StudentLayout';

export type CourseViewParams = {
  courseCode?: string;
};

const router = createBrowserRouter([
  {
    path: '/',
    element: <LoginPage />,
  },
  {
    path: '/logged-in',
    element: <LoggedInPage />,
  },
  {
    path: '/',
    element: <AuthLayout />,
    children: [
      {
        path: '/admin',
        element: <AdminLayout />,
        children: [
          {
            // Return a spinner, because it should be redirect to the first available course
            path: '/admin/:courseCode',
            element: <FullSpinner />,
          },
          {
            path: '/admin/:courseCode/students',
            element: <AdminStudentsPage />,
          },
          {
            path: '/admin/:courseCode/groups/new',
            element: <AddGroupsPage />,
          },
          {
            path: '/admin/:courseCode/groups',
            element: <GroupsPage />,
          },
          {
            // Return a spinner, because it should be redirect to the first available course
            path: '/admin',
            element: <FullSpinner />,
          },
        ],
      },
      {
        path: '/',
        element: <StudentLayout />,
        children: [
          {
            path: '/questions',
            element: <QuestionPage />,
          },
          {
            path: '/report',
            element: <StudentReportPage />,
          },
        ],
      },
    ],
  },
]);

export default router;
