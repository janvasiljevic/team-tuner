/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * TT API
 * Team tuner API documentation
 * OpenAPI spec version: 1.0
 */
import type { OutBigFiveBoxPlot } from './outBigFiveBoxPlot';
import type { OutStudentOut } from './outStudentOut';

export interface OutDetailedGroupOut {
  bigFiveBoxPlot: OutBigFiveBoxPlot;
  courseCode: string;
  courseColour: string;
  courseName: string;
  createdAt: string;
  id: string;
  name: string;
  students: OutStudentOut[];
  updateAt: string;
}
