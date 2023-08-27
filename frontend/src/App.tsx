import { MantineProvider, useMantineTheme } from '@mantine/core';
import { ModalsProvider } from '@mantine/modals';
import { Notifications, showNotification } from '@mantine/notifications';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Suspense } from 'react';
import { RouterProvider } from 'react-router-dom';

import router from './router';

import { useMediaQuery } from '@mantine/hooks';
import { NavigationProgress } from '@mantine/nprogress';
import { IconBug } from '@tabler/icons-react';

import '@fontsource-variable/raleway';
import '@fontsource/libre-baskerville';
import '@fontsource-variable/fira-code';
import AddCourseModal from './modals/AddCourseModal';

const queryClient = new QueryClient({
  defaultOptions: {
    mutations: {
      onSettled(data, error) {
        const axError = error as AxiosError;

        if (!axError) return;
        if (!axError.response) return;

        if (axError.response?.status === 403) {
          showNotification({
            title: 'Unauthorized',
            message: 'You are not authorized to perform this action',
            icon: <IconBug />,
          });
        }
      },
    },
    queries: {
      retry: (failureCount, error) => {
        const axError = error as AxiosError;

        // Dont retry on the following errors
        switch (axError.response?.status) {
          case 401: // Unauthorized
          case 403: // Forbidden
          case 404: // Not found
            return false;
        }

        if (failureCount < 2) return true;

        return false;
      },
    },
  },
});

export const modalsList = {
  addCourse: AddCourseModal,
} as const;

function App() {
  const theme = useMantineTheme();
  // theme.fn.smallerThan("md") returns something like @media (max-width: 61.9375em)
  // useMediaQuery accepts onlu the (max-width: 61.9375em) part
  const smallerThanSm = useMediaQuery(
    theme.fn.smallerThan('sm').replace('@media ', ''),
  );

  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider
        withGlobalStyles
        withNormalizeCSS
        theme={{
          // fontFamily: 'Libre Baskerville, serif',
          // fontFamilyMonospace: 'Fira Code, monospace',
          // defaultRadius: 'sm',
          // headings: {
          //   fontFamily: 'Ralway, sans-serif',
          //   fontWeight: 900,
          // },
          fontFamily: 'Roboto, sans-serif',
          headings: {
            fontFamily: 'Roboto, sans-serif',
          },
        }}
      >
        <ModalsProvider
          modals={modalsList}
          // We want to use the full screen modal on mobile devices
          modalProps={{
            fullScreen: smallerThanSm,
            zIndex: 3,
          }}
        >
          <NavigationProgress />
          <Notifications zIndex={4} position="top-left" />
          <Suspense fallback={<></>}>
            <RouterProvider router={router} />
          </Suspense>
        </ModalsProvider>
      </MantineProvider>
    </QueryClientProvider>
  );
}

declare module '@mantine/modals' {
  export interface MantineModalsOverride {
    modals: typeof modalsList;
  }
}

export default App;
