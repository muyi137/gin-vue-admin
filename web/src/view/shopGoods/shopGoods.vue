<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="商品名称">
          <el-input v-model="searchInfo.goodsName" placeholder="搜索条件" />
        </el-form-item>
        <!-- <el-form-item label="规格">
          <el-input v-model="searchInfo.attrName" placeholder="规格" />
        </el-form-item> -->
        <!-- <el-form-item label="批准字号">
          <el-input v-model="searchInfo.approvalNumber" placeholder="搜索条件" />
        </el-form-item> -->
        <el-form-item label="厂家">
          <el-input v-model="searchInfo.productionName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column fixed="right" align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updateShopGoodsFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="edit" size="small" title="{{ filterDict(parseInt(scope.row.dataSources),mallNameOptions) }}/{{ formatDate(scope.row.CreatedAt,'yyyy-MM-dd hh:mm:ss') }}" class="table-button" @click="openDrawer(scope.row)">
              查看价格</el-button>
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品id" prop="goodsId" width="120" />
        <el-table-column fixed align="left" label="图片地址" width="120">
          <template #default="scope">
            <img :src="scope.row.goodsImage" style="width: 100%; height: 100%;">
          </template>
        </el-table-column>
        <el-table-column fixed align="left" label="商品名称" sortable prop="goodsName" width="140" />
        <!-- <el-table-column align="left" label="分类id" prop="categoryId" width="120" /> -->
        <!-- <el-table-column align="left" label="类型" prop="scopeId" width="120" /> -->
        <!-- <el-table-column align="left" label="规格" sortable prop="attrName"  /> -->
        <el-table-column label="规格-单位">
          <template #default="scope">
            {{ scope.row.attrName }}-{{ scope.row.packageUnit }}
          </template>

        </el-table-column>
        <!-- <el-table-column align="left" label="单位" prop="packageUnit"  /> -->
        <!-- <el-table-column align="left" label="包装数量" prop="packageNum"  /> -->
        <el-table-column align="left" label="中包装" sortable prop="middlePackage" />
        <el-table-column align="left" label="厂家" sortable prop="productionName" width="180" />
        <!-- <el-table-column align="left" label="剂型" prop="dosage" width="120" /> -->
        <!-- <el-table-column align="left" label="近效期"   width="120"  >
          <template #default="scope">
             {{ formatTimeToStr(scope.row.expireTime,'yyyy-MM-dd') }}
          </template>
        </el-table-column> -->


        <!-- <el-table-column align="left" label="销售价" sortable prop="salePrice"  />
        <el-table-column align="left" label="促销价" sortable prop="promotionPrice"  /> -->
        <!-- <el-table-column align="left" label="最高促销价" prop="maxPromotionPrice" /> -->
        <!-- <el-table-column align="left" label="立省价格" prop="maxDiscountPrice"  /> -->
        <!-- <el-table-column align="left" label="批准字号" prop="approvalNumber" width="120" /> -->

        <!-- <el-table-column align="left" label="库存量" prop="stockNum"  /> -->

        <!-- <el-table-column align="left" label="促销状态" prop="promotionStatus" /> -->
        <el-table-column align="left" label="更新日期" sortable width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt,'yyyy-MM-dd') }}
          <!-- {{ formatTimeToStr(scope.row.UpdatedAt, 'yyyy-MM-dd HH:mm:ss') }} -->
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="商品编号" prop="goodsCode" width="120" /> -->
        <!-- <el-table-column align="left" label="erp code" prop="erpCode" width="120" /> -->
        <!-- <el-table-column align="left" label="数据来源" prop="dataSources" width="120" >
           <template #default="scope">
            {{ filterDict(parseInt(scope.row.dataSources),mallNameOptions) }}
            </template>
        </el-table-column> -->


      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 50, 200, 1000]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="商品id:" prop="goodsId">
          <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品名称:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类id:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="scopeId">
          <el-input v-model="formData.scopeId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="规格:" prop="attrName">
          <el-input v-model="formData.attrName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单位:" prop="packageUnit">
          <el-input v-model="formData.packageUnit" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="剂型:" prop="dosage">
          <el-input v-model="formData.dosage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批准字号:" prop="approvalNumber">
          <el-input v-model="formData.approvalNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="近效期:" prop="expireTime">
          <el-date-picker v-model="formData.expireTime" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
        <el-form-item label="包装数量:" prop="packageNum">
          <el-input v-model.number="formData.packageNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售价格:" prop="salePrice">
          <el-input-number v-model="formData.salePrice" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="促销价格:" prop="promotionPrice">
          <el-input-number v-model="formData.promotionPrice" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="最高促销价格:" prop="maxPromotionPrice">
          <el-input-number v-model="formData.maxPromotionPrice" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="立省价格 =原价-促销价:" prop="maxDiscountPrice">
          <el-input-number v-model="formData.maxDiscountPrice" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="中包装:" prop="middlePackage">
          <el-input v-model.number="formData.middlePackage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存量:" prop="stockNum">
          <el-input v-model.number="formData.stockNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="厂家:" prop="productionName">
          <el-input v-model="formData.productionName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="促销状态:" prop="promotionStatus">
          <el-input v-model="formData.promotionStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片地址:" prop="goodsImage">
          <el-input v-model="formData.goodsImage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品编号:" prop="goodsCode">
          <el-input v-model="formData.goodsCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="erp code:" prop="erpCode">
          <el-input v-model="formData.erpCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据来源:" prop="dataSources">
          <el-input v-model="formData.dataSources" :clearable="true" placeholder="请输入" />
        </el-form-item>

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer v-model="drawer" @open="loadDrawerData" size="70%"
      :title="activeRow.goodsName + '/' + activeRow.attrName + '/' + activeRow.productionName">
  
     <el-form   label-width="120px">
       <el-form-item label="筛选">
          <el-checkbox-group v-model="checkList" @change="checkboxTap">
                  <el-checkbox v-for="city in cities" :key="city" :label="city">{{  city }}</el-checkbox> 
          </el-checkbox-group>
        </el-form-item>
      </el-form>
   
      <el-table :data="tableDrawerData" v-loading="loading"  >
        <el-table-column prop="salePrice" sortable label="价格" >
          <template #default="scope">
            <span>{{ scope.row.salePrice }}/{{ scope.row.promotionPrice }}</span>
          </template>

        </el-table-column>
        <el-table-column prop="maxPromotionPrice" sortable label="促销" />
        <el-table-column prop="goodsName" label="商品名称"   width="200" >
        <template #default="scope" >
           
          <el-link type="primary" :underline="false" @click="goodsNameTap(scope.row)">{{ scope.row.goodsName }}</el-link>
         
        </template>
        </el-table-column>
        <el-table-column prop="attrName" label="规格" sortable width="100"/>
        <el-table-column prop="expireTime" label="效期" sortable>
          <template #default="scope">
            {{ formatTimeToStr(scope.row.expireTime, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column prop="productionName" label="厂家" width="250" />
        <el-table-column prop="dataSources" label="来源" sortable width="120" >
          <template #default="scope">
            {{ filterDict(parseInt(scope.row.dataSources),mallNameOptions) }}
            </template>
          </el-table-column>
        
         <el-table-column prop="tag" label="标签" sortable >
          <template #default="scope">
            {{ scope.row.tag }}
            </template>
          </el-table-column>
 
      <el-table-column   label="库存量" prop="stockNum"  >
        <template #default="scope">
          约{{ scope.row.stockNum }} {{ scope.row.stockNum }}
          </template>
      </el-table-column>
        <el-table-column   label="更新日期" sortable width="180">
          <template #default="scope">
          {{ formatTimeToStr(scope.row.CreatedAt, 'yyyy-MM-dd hh:mm') }}
           
          <!-- {{ formatTimeToStr(scope.row.UpdatedAt, 'yyyy-MM-dd HH:mm:ss') }} -->
          </template>
        </el-table-column>

        <el-table-column prop="goodsUrl" label="操作" >
          <template #default="scope">
          <el-link    :href="scope.row.goodsUrl" target="_blank">
            <i class="el-icon-link"></i>
             <el-button type="danger" round>加购</el-button>
             
          </el-link>

          <el-link    :href="scope.row.goodsUrl" target="_blank">
            <i class="el-icon-link"></i>
             <el-button type="primary" round>查看</el-button>
             
          </el-link>
          </template>
           
          </el-table-column>
      </el-table>
     
       
 
       
    </el-drawer>

  </div>


</template>
  
<script>
export default {
  name: 'ShopGoods'
}
</script>

<script setup>
import {
  createShopGoods,
  deleteShopGoods,
  deleteShopGoodsByIds,
  updateShopGoods,
  findShopGoods,
  getShopGoodsList
} from '@/api/shopGoods'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { formatTimeToStr } from '@/utils/date'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
 
const loading = ref(true)

// 自动化生成的字典（可能为空）以及字段
const mallNameOptions = ref([])
const formData = ref({
  goodsId: 0,
  goodsName: '',
  categoryId: 0,
  scopeId: '',
  attrName: '',
  packageUnit: '',
  dosage: '',
  approvalNumber: '',
  expireTime: new Date(),
  packageNum: 0,
  salePrice: 0,
  promotionPrice: 0,
  maxPromotionPrice: 0,
  maxDiscountPrice: 0,
  middlePackage: 0,
  stockNum: 0,
  productionName: '',
  promotionStatus: '',
  goodsImage: '',
  goodsCode: '',
  erpCode: '',
  dataSources: '',
  addtime: '',
})


// 验证规则
const rule = reactive({
  goodsName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 50
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

 

// 查询
const getTableData = async () => {
  const table = await getShopGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  mallNameOptions.value = await getDictFunc('mallName')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 打开抽屉查看价格
const drawer = ref(false)
const activeRow = ref({})
const tableDrawerData = ref([])
const searchDrawerInfo = ref({})
const checkList = ref([])
const cities = ref({})

const goodsNameTap = (row) => {
  activeRow.value = row

  const a = row.attrFilter
  // eslint-disable-next-line no-const-assign
  cities.value = a.split(',').filter(function(s) {
    return s && s.trim() // 去空格
  })
 

  searchDrawerInfo.value = { 
    // attrName: row.attrName,
    productionName: NameFilter(row.productionName),
  } 
  searchDrawerInfo.value.goodsName = NameFilter(row.goodsName)

}

const openDrawer = (row) => {
  checkList.value = []
  drawer.value = true
  activeRow.value = row

  // cities = row.attrFilter.split(',')
  const a = row.attrFilter
  // eslint-disable-next-line no-const-assign
  cities.value = a.split(',').filter(function(s) {
    return s && s.trim() // 去空格
  })
  
  searchDrawerInfo.value = { 
    // attrName: row.attrName,
    productionName: NameFilter(row.productionName),
  }

  searchDrawerInfo.value.goodsName = NameFilter(row.goodsName)

  loadDrawerData()

  console.log(searchDrawerInfo.value)
}

//过滤字符串 替换（为(
const NameFilter = (str) => {

  if ((str.indexOf('（') || str.indexOf('(')) > -1) {
    str = str.substring(0, str.replace('（', '(').indexOf('('))
  }

  const index = str.lastIndexOf('(')
  if (index > -1) {
    str = str.substring(0, index)
  }
  return str
  console.log(str)
}

const loadDrawerData = async () => {
  const tableDrawer = await getShopGoodsList({ pageSize: 99999, ...searchDrawerInfo.value })
  if (tableDrawer.code === 0) {
    tableDrawerData.value = tableDrawer.data.list
  }
  loading.value = false
}

const checkboxTap = async(val) => {
  console.log(val)
 
  const tableDrawer = await getShopGoodsList({ pageSize: 99999, attrFilter: val.join(','), ...searchDrawerInfo.value })
  if (tableDrawer.code === 0) {
    tableDrawerData.value = tableDrawer.data.list
  }
}
 
// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteShopGoodsFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteShopGoodsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateShopGoodsFunc = async (row) => {
  const res = await findShopGoods({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reshopGoods
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteShopGoodsFunc = async (row) => {
  const res = await deleteShopGoods({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    goodsId: 0,
    goodsName: '',
    categoryId: 0,
    scopeId: '',
    attrName: '',
    packageUnit: '',
    dosage: '',
    approvalNumber: '',
    expireTime: new Date(),
    packageNum: 0,
    salePrice: 0,
    promotionPrice: 0,
    maxPromotionPrice: 0,
    maxDiscountPrice: 0,
    middlePackage: 0,
    stockNum: 0,
    productionName: '',
    promotionStatus: '',
    goodsImage: '',
    goodsCode: '',
    erpCode: '',
    dataSources: '',
    addtime: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createShopGoods(formData.value)
        break
      case 'update':
        res = await updateShopGoods(formData.value)
        break
      default:
        res = await createShopGoods(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
