<template>
  <div class="lmstudio-settings">
    <div class="section-header">
      <h2>{{ $t('lmStudioSettings.title') }}</h2>
      <p class="section-description">{{ $t('lmStudioSettings.description') }}</p>
    </div>

    <div class="settings-group">
      <!-- LM Studio 服务状态 -->
      <div class="setting-row">
        <div class="setting-info">
          <label>{{ $t('lmStudioSettings.status.label') }}</label>
          <p class="desc">{{ $t('lmStudioSettings.status.desc') }}</p>
        </div>
        <div class="setting-control">
          <div class="status-display">
            <t-tag 
              v-if="testing"
              theme="default"
              variant="light"
            >
              <t-icon name="loading" class="status-icon spinning" />
              {{ $t('lmStudioSettings.status.testing') }}
            </t-tag>
            <t-tag 
              v-else-if="connectionStatus === true"
              theme="success"
              variant="light"
            >
              <t-icon name="check-circle-filled" />
              {{ $t('lmStudioSettings.status.available') }}
            </t-tag>
            <t-tag 
              v-else-if="connectionStatus === false"
              theme="danger"
              variant="light"
            >
              <t-icon name="close-circle-filled" />
              {{ $t('lmStudioSettings.status.unavailable') }}
            </t-tag>
            <t-tag 
              v-else
              theme="default"
              variant="light"
            >
              <t-icon name="help-circle" />
              {{ $t('lmStudioSettings.status.untested') }}
            </t-tag>
            <t-button 
              size="small" 
              variant="outline"
              :loading="testing"
              @click="testConnection"
            >
              <t-icon name="refresh" />
              {{ $t('lmStudioSettings.status.retest') }}
            </t-button>
          </div>
        </div>
      </div>

      <!-- LM Studio 服务地址 -->
      <div class="setting-row">
        <div class="setting-info">
          <label>{{ $t('lmStudioSettings.address.label') }}</label>
          <p class="desc">{{ $t('lmStudioSettings.address.desc') }}</p>
        </div>
        <div class="setting-control">
          <div class="url-control-group">
            <t-input 
              v-model="localBaseUrl" 
              :placeholder="$t('lmStudioSettings.address.placeholder')"
              disabled
              style="flex: 1;"
            />
          </div>
          <t-alert 
            v-if="connectionStatus === false"
            theme="warning"
            :message="$t('lmStudioSettings.address.failed')"
            style="margin-top: 8px;"
          />
        </div>
      </div>

    </div>

    <!-- 可用的模型 -->
    <div v-if="connectionStatus === true" class="model-category-section">
      <div class="category-header">
        <div class="header-info">
          <h3>{{ $t('lmStudioSettings.installed.title') }}</h3>
          <p>{{ $t('lmStudioSettings.installed.desc') }}</p>
        </div>
        <t-button 
          size="small" 
          variant="text"
          :loading="loadingModels"
          @click="refreshModels"
        >
          <t-icon name="refresh" />{{ $t('common.refresh') }}
        </t-button>
      </div>
      
      <div v-if="loadingModels" class="loading-state">
        <t-loading size="small" />
        <span>{{ $t('common.loading') }}</span>
      </div>
      <div v-else-if="availableModels.length > 0" class="model-list-container">
        <div v-for="model in availableModels" :key="model.id" class="model-card">
          <div class="model-info">
            <div class="model-name">{{ model.id }}</div>
            <div class="model-meta">
              <span class="model-owner">{{ $t('lmStudioSettings.installed.ownedBy') }}: {{ model.owned_by }}</span>
            </div>
          </div>
        </div>
      </div>
      <div v-else class="empty-state">
        <p class="empty-text">{{ $t('lmStudioSettings.installed.empty') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useSettingsStore } from '@/stores/settings'
import { MessagePlugin } from 'tdesign-vue-next'
import { useI18n } from 'vue-i18n'
import { checkLMStudioStatus, listLMStudioModels, type LMStudioModelInfo } from '@/api/initialization'

const settingsStore = useSettingsStore()
const { t } = useI18n()

const localBaseUrl = ref(settingsStore.settings.lmStudioConfig?.baseUrl ?? '')

const testing = ref(false)
const connectionStatus = ref<boolean | null>(null)
const loadingModels = ref(false)
const availableModels = ref<LMStudioModelInfo[]>([])

// 测试连接
const testConnection = async () => {
  testing.value = true
  connectionStatus.value = null
  
  try {
    // 保存配置
    settingsStore.updateLMStudioConfig({ baseUrl: localBaseUrl.value })
    
    // 调用真实 LM Studio API 测试连接
    const result = await checkLMStudioStatus()
    
    // 如果接口返回了 baseUrl 且与当前输入框的值不同，更新为接口返回的值
    if (result.baseUrl && result.baseUrl !== localBaseUrl.value) {
      localBaseUrl.value = result.baseUrl
      settingsStore.updateLMStudioConfig({ baseUrl: result.baseUrl })
    }
    
    connectionStatus.value = result.available
    
    if (connectionStatus.value) {
      MessagePlugin.success(t('lmStudioSettings.toasts.connected'))
      refreshModels()
    } else {
      MessagePlugin.error(result.error || t('lmStudioSettings.toasts.connectFailed'))
    }
  } catch (error: any) {
    connectionStatus.value = false
    MessagePlugin.error(error.message || t('lmStudioSettings.toasts.connectFailed'))
  } finally {
    testing.value = false
  }
}

// 刷新模型列表
const refreshModels = async () => {
  loadingModels.value = true
  
  try {
    // 调用真实 LM Studio API 获取模型列表
    const models = await listLMStudioModels()
    availableModels.value = models
  } catch (error: any) {
    console.error('获取模型列表失败:', error)
    MessagePlugin.error(error.message || t('lmStudioSettings.toasts.listFailed'))
  } finally {
    loadingModels.value = false
  }
}

// 初始化 LM Studio 服务地址
const initLMStudioBaseUrl = async () => {
  try {
    const result = await checkLMStudioStatus()
    // 如果接口返回了 baseUrl，优先使用接口返回的值
    if (result.baseUrl) {
      localBaseUrl.value = result.baseUrl
      // 如果 store 中没有保存过，也保存到 store 中
      if (!settingsStore.settings.lmStudioConfig?.baseUrl) {
        settingsStore.updateLMStudioConfig({ baseUrl: result.baseUrl })
      }
    } else if (!localBaseUrl.value) {
      // 如果接口没返回且 store 中也没有，使用默认值
      localBaseUrl.value = 'http://localhost:1234/v1'
    }
    
    // 直接使用初始化时获取的状态，避免重复调用
    connectionStatus.value = result.available
    if (result.available) {
      refreshModels()
    }
    
    return result
  } catch (error) {
    console.error('初始化 LM Studio 地址失败:', error)
    // 如果获取失败，使用默认值或 store 中的值
    if (!localBaseUrl.value) {
      localBaseUrl.value = 'http://localhost:1234/v1'
    }
    return null
  }
}

// 组件挂载时自动检查连接
onMounted(async () => {
  // 初始化服务地址，如果启用则直接使用返回的状态，避免重复调用
  await initLMStudioBaseUrl()
})
</script>

<style lang="less" scoped>
.lmstudio-settings {
  width: 100%;
}

.section-header {
  margin-bottom: 32px;

  h2 {
    font-size: 20px;
    font-weight: 600;
    color: #333333;
    margin: 0 0 8px 0;
  }

  .section-description {
    font-size: 14px;
    color: #666666;
    margin: 0;
    line-height: 1.5;
  }
}

.settings-group {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.setting-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 20px 0;
  border-bottom: 1px solid #e5e7eb;

  &:last-child {
    border-bottom: none;
  }
}

.setting-info {
  flex: 1;
  padding-right: 32px;

  label {
    font-size: 15px;
    font-weight: 500;
    color: #333333;
    display: block;
    margin-bottom: 4px;
  }

  .desc {
    font-size: 13px;
    color: #666666;
    margin: 0;
    line-height: 1.6;
  }
}

.setting-control {
  flex-shrink: 0;
  min-width: 360px;
  max-width: 360px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.status-display {
  display: flex;
  align-items: center;
  gap: 12px;

  .status-icon.spinning {
    animation: spin 1s linear infinite;
  }
}

.url-control-group {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 8px;
}

.model-category-section {
  margin-top: 32px;
  margin-bottom: 32px;
  padding-top: 32px;
  border-top: 1px solid #e5e7eb;

  &:first-of-type {
    margin-top: 24px;
    padding-top: 24px;
  }

  &:last-child {
    margin-bottom: 0;
  }
}

.category-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 24px;

  .header-info {
    flex: 1;

    h3 {
      font-size: 17px;
      font-weight: 600;
      color: #333333;
      margin: 0 0 6px 0;
    }

    p {
      font-size: 13px;
      color: #999999;
      margin: 0;
      line-height: 1.5;
    }
  }
}

.loading-state {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 48px 0;
  color: #999999;
  font-size: 14px;
}

.model-list-container {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;

  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
}

.model-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #fafafa;
  transition: all 0.2s;

  &:hover {
    border-color: #07C05F;
    background: #ffffff;
  }
}

.model-info {
  flex: 1;
  min-width: 0;

  .model-name {
    font-size: 14px;
    font-weight: 500;
    color: #333333;
    margin-bottom: 4px;
    font-family: monospace;
  }

  .model-meta {
    display: flex;
    gap: 12px;
    font-size: 12px;
    color: #666666;
  }
}

.empty-state {
  padding: 48px 0;
  text-align: center;

  .empty-text {
    font-size: 14px;
    color: #999999;
    margin: 0;
  }
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>

