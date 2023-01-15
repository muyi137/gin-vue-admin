<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
          <el-date-picker v-model="formData.expireTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="包装数量:" prop="packageNum">
          <el-input v-model.number="formData.packageNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售价格:" prop="salePrice">
          <el-input-number v-model="formData.salePrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="促销价格:" prop="promotionPrice">
          <el-input-number v-model="formData.promotionPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="最高促销价格:" prop="maxPromotionPrice">
          <el-input-number v-model="formData.maxPromotionPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="立省价格 =原价-促销价:" prop="maxDiscountPrice">
          <el-input-number v-model="formData.maxDiscountPrice" :precision="2" :clearable="true"></el-input-number>
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
        <el-form-item label="添加时间:" prop="addtime">
          <el-input v-model="formData.addtime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateShopGoods,
  findShopGoods
} from '@/api/shopGoods'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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
               goodsName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findShopGoods({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reshopGoods
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
