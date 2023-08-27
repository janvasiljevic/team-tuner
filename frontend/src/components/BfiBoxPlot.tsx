import { BoxPlotDatum, ResponsiveBoxPlot } from '@nivo/boxplot';
import { useMemo } from 'react';
import { OutBigFiveBoxPlot, OutBoxPlotItem } from '../api/model';

type Props = {
  serverData?: OutBigFiveBoxPlot;
  showLabels?: boolean;
};

const BfiBoxPlot = ({ serverData, showLabels = true }: Props) => {
  const data = useMemo(() => {
    if (!serverData) return [];

    const data: BoxPlotDatum[] = [];

    const addToData = (name: string, item: OutBoxPlotItem) => {
      const newData = item.dataPoints.map((d) => ({
        group: name,
        subgroup: 0,
        mu: 1,
        sd: 1,
        n: 1,
        value: d,
      }));

      data.push(...newData);
    };

    addToData('Openness', serverData.openness);
    addToData('Conscientiousness', serverData.conscientious);
    addToData('Extraversion', serverData.extraversion);
    addToData('Agreeableness', serverData.agreeableness);
    addToData('Neuroticism', serverData.neuroticism);

    return data;
  }, [serverData]);

  return (
    <ResponsiveBoxPlot
      data={data}
      margin={{ top: 40, right: 40, bottom: 50, left: 40 }}
      minValue={0}
      maxValue={1}
      subGroupBy="subgroup"
      colorBy="group"
      padding={0.5}
      enableGridY={true}
      enableGridX={false}
      colors={{ scheme: 'pastel2' }}
      axisLeft={{
        tickValues: 3,
      }}
      axisBottom={showLabels ? {} : null}
      borderRadius={2}
      borderWidth={1}
      gridYValues={[0.25, 0.5, 0.75]}
      borderColor={{
        from: 'color',
        modifiers: [['darker', 0.3]],
      }}
      medianWidth={1}
      medianColor={{
        from: 'color',
        modifiers: [['darker', 0.3]],
      }}
      whiskerEndSize={0.2}
      whiskerColor={{
        from: 'color',
        modifiers: [['darker', 0.3]],
      }}
      motionConfig="stiff"
    />
  );
};

export default BfiBoxPlot;
