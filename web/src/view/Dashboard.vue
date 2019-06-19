<template>
    <div>
        <el-card>
            <el-table
                    class="app-table"
                    size="medium"
                    v-loading="tableLoading"
                    :data="tableData">
                <!--<el-table-column type="selection" width="55"></el-table-column>-->
                <el-table-column prop="status" label="状态" width="100">
                    <template slot-scope="scope">
                        <i class="iconfont green icon-cloud"></i>
                    </template>
                </el-table-column>
                <el-table-column prop="metadata.name" :label="$t('name')"></el-table-column>
                <el-table-column prop="status.replicas" label="副本数量"></el-table-column>
                <el-table-column prop="spec.strategy.type" label="流量百分比">
                    <template slot-scope="scope">
                        <span>100%</span>
                    </template>
                </el-table-column>
                <el-table-column prop="metadata.creationTimestamp" label="创建时间" :formatter="dateFmt"></el-table-column>
                <el-table-column prop="status" label="操作" width="80">
                    <template slot-scope="scope">
                        <a class="app-link" href="#"><i class="iconfont blue icon-more-v"></i></a>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <div class="app-center-info">监控信息</div>
        <el-card v-for="n in nodes" :key="n.nodeName">
            <p>节点 {{n.nodeName}}</p>
            <div class="app-charts-wrap">
                <div>
                    <el-progress width="200" type="circle" :percentage="n.memory"></el-progress>
                    <p>Memory</p>
                </div>
                <div>
                    <el-progress width="200" type="circle" :percentage="n.cpu"></el-progress>
                    <p>CPU</p>
                </div>
            </div>
        </el-card>
    </div>
</template>

<script>
    import { k8sDeployments, k8sResources } from '@/api/k8s'
    export default {
        data() {
            return {
                tableLoading: false,
                tableData: [],
                percentageCPU: 0,
                percentageMemory: 0,
                nodes: []
            }
        },
        methods: {
            loadDeployments() {
                this.tableLoading = true;
                k8sDeployments({offset: 0, limit : 999}).then(res => {
                    if (res.items) {
                        this.tableData = res.items;
                        this.tableLoading = false;
                    }
                })
            },
            loadResources() {
                this.tableLoading = true;
                k8sResources({offset: 0, limit : 999}).then(res => {
                    console.log(res)
                    if (res.length > 0) {
                        this.nodes = res;
                        this.tableLoading = false;
                    }
                })
            },
            dateFmt(row, column, data) {
                return data.replace("T", ' ').replace('Z', ' ')
            },
        },
        mounted() {
            this.loadDeployments();
            this.loadResources();
        },
    }
</script>