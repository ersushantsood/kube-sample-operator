package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ChaosPhase string

type ChaosStatus struct {
   Phase ChaosPhase `json:"phase"`
   Reason string `json:"reason,omitempty"`
   Experiment ExperimentStatus `json:"experiment"`
}

type ExperimentPhase string

const (
    ExperimentPhaseRunning ExperimentPhase = "Running"
    ExperimentPhaseFailed  ExperimentPhase = "Failed"
    ExperimentPhaseFinished ExperimentPhase = "Finished"
)

type ExperimentStatus struct {
    Phase ExperimentPhase `json:"phase,omitempty"`
    Reason string `json:"reason,omitempty"`
    StartTime *metav1.Time `json:"startTime,omitempty"`
    EndTime *metav1.Time `json:"endTime,omitempty"`
    Pods []ChaosStatus `json:"podChaos,omitempty"`
}