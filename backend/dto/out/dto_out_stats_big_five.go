package out

import (
	"jv/team-tone-tuner/model"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

type BoxPlotItem struct {
	Min        float64   `json:"min" validate:"required"`
	Q1         float64   `json:"q1" validate:"required"`
	Q3         float64   `json:"q3" validate:"required"`
	Max        float64   `json:"max" validate:"required"`
	Mean       float64   `json:"mean" validate:"required"`
	DataPoints []float64 `json:"dataPoints" validate:"required"`
}

type BigFiveBoxPlot struct {
	Openness      BoxPlotItem `json:"openness" validate:"required"`
	Conscientious BoxPlotItem `json:"conscientious" validate:"required"`
	Extraversion  BoxPlotItem `json:"extraversion" validate:"required"`
	Agreeableness BoxPlotItem `json:"agreeableness" validate:"required"`
	Neuroticism   BoxPlotItem `json:"neuroticism" validate:"required"`
}

func NewBigFiveBoxPlotFromReports(reports []*model.BfiReport) BigFiveBoxPlot {
	openness := make([]float64, len(reports))
	conscientious := make([]float64, len(reports))
	extraversion := make([]float64, len(reports))
	agreeableness := make([]float64, len(reports))
	neuroticism := make([]float64, len(reports))

	for i, report := range reports {
		openness[i] = report.Openness.PointsNormalized
		conscientious[i] = report.Conscientiousness.PointsNormalized
		extraversion[i] = report.Extraversion.PointsNormalized
		agreeableness[i] = report.Agreeableness.PointsNormalized
		neuroticism[i] = report.Neuroticism.PointsNormalized
	}

	return BigFiveBoxPlot{
		Openness:      boxPlotItemFromSeries(openness),
		Conscientious: boxPlotItemFromSeries(conscientious),
		Extraversion:  boxPlotItemFromSeries(extraversion),
		Agreeableness: boxPlotItemFromSeries(agreeableness),
		Neuroticism:   boxPlotItemFromSeries(neuroticism),
	}
}

func boxPlotItemFromSeries(series []float64) BoxPlotItem {

	if len(series) == 0 {
		return BoxPlotItem{
			DataPoints: series,
		}
	}

	sort.Float64s(series)

	n := len(series)
	min := series[0]
	q1 := series[int(math.Round(float64(n)/4))-1]
	q3 := series[int(math.Round(float64(3*n)/4))-1]
	max := series[n-1]
	mean := stat.Mean(series, nil)

	return BoxPlotItem{
		Min:        min,
		Q1:         q1,
		Q3:         q3,
		Max:        max,
		Mean:       mean,
		DataPoints: series,
	}
}
