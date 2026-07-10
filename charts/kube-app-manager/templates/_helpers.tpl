{{- define "kube-app-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "kube-app-manager.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := include "kube-app-manager.name" . -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "kube-app-manager.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
{{- default (include "kube-app-manager.fullname" .) .Values.serviceAccount.name -}}
{{- else -}}
{{- default "default" .Values.serviceAccount.name -}}
{{- end -}}
{{- end -}}

{{- define "kube-app-manager.tlsSecretName" -}}
{{- if .Values.ingressRoute.tls.secretName -}}
{{- .Values.ingressRoute.tls.secretName -}}
{{- else -}}
{{- printf "%s-tls" (include "kube-app-manager.fullname" .) -}}
{{- end -}}
{{- end -}}
