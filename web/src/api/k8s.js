import {get, post} from '@/lib/fetch.js'

export function k8sDeployments(params) {
    return get('/k8s/deployments', params)
}

export function k8sResources(params) {
    return get('/k8s/resources', params)
}