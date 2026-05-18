# docker app： BASE_URL=http://localhost:18080/api/v1 ./scripts/api_test.sh
#!/usr/bin/env bash

set -e

BASE_URL="${BASE_URL:-http://localhost:8080/api/v1}"

echo "BASE_URL = $BASE_URL"
echo "开始接口测试..."

echo
echo "1. 健康检查"
curl -s "$BASE_URL/ping"
echo

echo
echo "2. 注册用户 tom"
curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username":"tom_test",
    "password":"123456",
    "nickname":"Tom Test"
  }'
echo

echo
echo "3. 注册用户 jack"
curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username":"jack_test",
    "password":"123456",
    "nickname":"Jack Test"
  }'
echo

echo
echo "4. tom 登录并提取 token"
TOM_LOGIN_RESP=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username":"tom_test",
    "password":"123456"
  }')

echo "$TOM_LOGIN_RESP"

TOM_TOKEN=$(echo "$TOM_LOGIN_RESP" | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ -z "$TOM_TOKEN" ]; then
  echo "提取 TOM_TOKEN 失败"
  exit 1
fi

echo "TOM_TOKEN 提取成功"

echo
echo "5. jack 登录并提取 token"
JACK_LOGIN_RESP=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username":"jack_test",
    "password":"123456"
  }')

echo "$JACK_LOGIN_RESP"

JACK_TOKEN=$(echo "$JACK_LOGIN_RESP" | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ -z "$JACK_TOKEN" ]; then
  echo "提取 JACK_TOKEN 失败"
  exit 1
fi

echo "JACK_TOKEN 提取成功"

echo
echo "6. 查看当前用户 me"
curl -s "$BASE_URL/users/me" \
  -H "Authorization: Bearer $TOM_TOKEN"
echo

echo
echo "7. jack 发布视频 1"
VIDEO1_RESP=$(curl -s -X POST "$BASE_URL/videos" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JACK_TOKEN" \
  -d '{
    "title":"jack测试视频1",
    "description":"接口测试视频1",
    "video_url":"https://example.com/video1.mp4",
    "cover_url":"https://example.com/cover1.jpg"
  }')

echo "$VIDEO1_RESP"

VIDEO1_ID=$(echo "$VIDEO1_RESP" | sed -n 's/.*"video_id":\([0-9]*\).*/\1/p')

if [ -z "$VIDEO1_ID" ]; then
  echo "提取 VIDEO1_ID 失败"
  exit 1
fi

echo "VIDEO1_ID = $VIDEO1_ID"

echo
echo "8. jack 发布视频 2"
VIDEO2_RESP=$(curl -s -X POST "$BASE_URL/videos" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JACK_TOKEN" \
  -d '{
    "title":"jack测试视频2",
    "description":"接口测试视频2",
    "video_url":"https://example.com/video2.mp4",
    "cover_url":"https://example.com/cover2.jpg"
  }')

echo "$VIDEO2_RESP"

VIDEO2_ID=$(echo "$VIDEO2_RESP" | sed -n 's/.*"video_id":\([0-9]*\).*/\1/p')

if [ -z "$VIDEO2_ID" ]; then
  echo "提取 VIDEO2_ID 失败"
  exit 1
fi

echo "VIDEO2_ID = $VIDEO2_ID"

echo
echo "9. 查看视频详情"
curl -s "$BASE_URL/videos/$VIDEO1_ID"
echo

echo
echo "10. tom 关注 jack，假设 jack_test 的 user_id 为 2 或注册返回值对应 ID"
echo "注意：这里需要根据你的数据库实际 jack user_id 调整 TARGET_USER_ID。"
TARGET_USER_ID="${TARGET_USER_ID:-2}"

curl -s -X POST "$BASE_URL/users/$TARGET_USER_ID/follow" \
  -H "Authorization: Bearer $TOM_TOKEN"
echo

echo
echo "11. 查看 tom 的关注 Feed"
curl -s "$BASE_URL/feed/following?limit=10" \
  -H "Authorization: Bearer $TOM_TOKEN"
echo

echo
echo "12. tom 点赞视频"
curl -s -X POST "$BASE_URL/videos/$VIDEO1_ID/like" \
  -H "Authorization: Bearer $TOM_TOKEN"
echo

echo
echo "13. tom 评论视频"
curl -s -X POST "$BASE_URL/videos/$VIDEO1_ID/comments" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOM_TOKEN" \
  -d '{
    "content":"接口测试评论"
  }'
echo

echo
echo "14. 查看评论列表"
curl -s "$BASE_URL/videos/$VIDEO1_ID/comments"
echo

echo
echo "15. 再次查看视频详情，验证 like_count/comment_count"
curl -s "$BASE_URL/videos/$VIDEO1_ID"
echo

echo
echo "接口测试完成。"