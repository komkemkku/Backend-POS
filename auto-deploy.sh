#!/bin/bash

# Auto Deploy Script สำหรับ Backend POS
# ใช้สำหรับ auto commit และ push ทุกครั้งที่มีการแก้ไข

echo "🚀 Starting auto-deploy for Backend POS..."

# ตรวจสอบว่ามีการเปลี่ยนแปลงหรือไม่
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "📝 Changes detected, preparing to commit..."
    
    # เพิ่มไฟล์ทั้งหมด
    git add -A
    
    # สร้าง commit message อัตโนมัติ
    TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')
    BRANCH=$(git branch --show-current)
    
    # ตรวจสอบว่ามีไฟล์อะไรเปลี่ยนแปลงบ้าง
    CHANGED_FILES=$(git diff --cached --name-only | head -5)
    
    COMMIT_MSG="🔄 Auto-deploy: Backend updates ($TIMESTAMP)

📁 Modified files:
$(echo "$CHANGED_FILES" | sed 's/^/- /')

🌿 Branch: $BRANCH
⏰ Timestamp: $TIMESTAMP"

    # Commit การเปลี่ยนแปลง
    git commit -m "$COMMIT_MSG"
    
    # Push ไปยัง origin
    echo "📤 Pushing to origin/$BRANCH..."
    git push origin $BRANCH
    
    if [ $? -eq 0 ]; then
        echo "✅ Auto-deploy completed successfully!"
        echo "🚀 Railway will automatically redeploy the backend"
        echo "🔗 Check deployment status at: https://railway.app"
    else
        echo "❌ Push failed! Please check your Git configuration"
        exit 1
    fi
else
    echo "ℹ️  No changes detected. Nothing to deploy."
fi

echo "✨ Auto-deploy script finished."
