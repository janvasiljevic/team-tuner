import React from 'react';
import { BarDatum, ResponsiveBar } from '@nivo/bar';
import { useGetCourseCourseIdStatsDailyActivity } from '../api/courses/courses';
import { useAppStore } from '../store';

const DailyActivity = () => {
  const { selectedCourse } = useAppStore();

  const { data } = useGetCourseCourseIdStatsDailyActivity(
    selectedCourse?.id || '',
    {},
    { query: { enabled: !!selectedCourse } },
  );

  const reshapedData: BarDatum[] = React.useMemo(() => {
    if (!data) return [];

    return data.activity.map(
      (e): BarDatum => ({
        id: e.day,
        value: e.count,
      }),
    );
  }, [data]);

  return (
    <ResponsiveBar
      data={reshapedData}
      margin={{ top: 50, right: 0, bottom: 50, left: 0 }}
      padding={0.5}
      valueScale={{ type: 'linear' }}
      indexScale={{ type: 'band', round: true }}
      colors={{ scheme: 'pastel1' }}
      axisLeft={null}
      gridYValues={[0]}
      axisTop={null}
      axisRight={null}
      labelSkipWidth={12}
      labelSkipHeight={12}
      labelTextColor={{
        from: 'color',
        modifiers: [['darker', 1.6]],
      }}
      axisBottom={{
        format: (v) => {
          const split = v.split('-');

          return `${split[2]}/${split[1]}`;
        },
      }}
    />
  );
};

export default DailyActivity;
