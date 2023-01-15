<!--
 * @Descripttion: 
 * @version: 
 * @Author: yangdi <muyi137@163.com>
 * @Date: 2022-08-01 17:04:04
 * @LastEditors: yangdi <muyi137@163.com>
 * @LastEditTime: 2022-08-05 15:22:03
-->

<template>
  <div class="upload">
    <div class="gva-table-box">

      <div class="gva-btn-list">
        <el-upload
          class="excel-btn"
          :action="`${path}/shopGoods/importGoodsExcel`"
          :headers="{'x-token':userStore.token ,'x-user-id':userStore.userInfo.ID}"
          :on-success="loadExcel"
          :show-file-list="false"
        >
          <el-button size="small" type="primary" icon="upload">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
        <el-button class="excel-btn" size="small" type="success" icon="download" @click="downloadExcelTemplate()">下载模板</el-button> 
        <!-- <el-button key="danger" type="danger" text bg >第一列商品名称,商品规格，商品厂家</el-button> -->
        <el-link type="danger" style="margin-left: 20px">注意列表顺序 商品名称,商品规格，商品厂家</el-link>
      </div>
      <el-table :data="tableData" row-key="ID">
      <el-table-column align="center" label="导入商品信息">
        <el-table-column align="left" fixed label="ID" min-width="100" prop="ID" />
        <el-table-column align="left"  fixed   label="商品名称" min-width="160" prop="name" />
        <el-table-column align="left" fixed   label="路由Path" min-width="160" prop="path" />
         
      </el-table-column  >
         <el-table-column align="center" label="匹配信息">
              <el-table-column align="left" label="父节点" min-width="90" prop="parentId" />
              <el-table-column align="left" label="排序" min-width="70" prop="sort" />
              <el-table-column align="left" label="文件路径" min-width="360" prop="component" />
         </el-table-column>
        
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'goodsExcel',
}
</script>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { exportExcel, loadExcelData } from '@/api/goodsExcel'
import { getMenuList } from '@/api/menu'
import { ref } from 'vue'
const path = ref(import.meta.env.VITE_BASE_API)

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])

// 查询
const getTableData = async(f = () => {}) => {
  // const table = await f({ page: page.value, pageSize: pageSize.value })
  // if (table.code === 0) {
  //   tableData.value = table.data.list
  //   total.value = table.data.total
  //   page.value = table.data.page
  //   pageSize.value = table.data.pageSize
  // }
}
getTableData(getMenuList)

const userStore = useUserStore()

const handleExcelExport = (fileName) => {
  if (!fileName || typeof fileName !== 'string') {
    fileName = 'ExcelExport.xlsx'
  }
  exportExcel(tableData.value, fileName)
}
const loadExcel = (e) => {
  console.log(e)
  getTableData(loadExcelData(e.data.fileName))
}
const downloadExcelTemplate = () => {
  downloadTemplate('ExcelTemplate.xlsx')
}

</script>

<style lang="scss" scoped>
.btn-list{
  display: flex;
  margin-bottom: 12px;
  justify-content: flex-end;

}
.excel-btn+.excel-btn{
  margin-left: 10px;
}
</style>
