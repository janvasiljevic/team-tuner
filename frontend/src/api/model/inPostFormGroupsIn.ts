/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * TT API
 * Team tuner API documentation
 * OpenAPI spec version: 1.0
 */

export interface InPostFormGroupsIn {
  /** To which course the groups belong */
  courseId: string;
  /** Array length must be at least 4
Each element must be greater than 0 and less than 10 */
  groupSizes: number[];
  /** General settings for SA */
  iterations: number;
  temperature: number;
  weightConscientiousness: number;
  weightExtraversion: number;
  weightNeuroticism: number;
  /** Weights for SA */
  weightSatisfaction: number;
}
