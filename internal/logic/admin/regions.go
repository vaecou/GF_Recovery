package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
	"sort"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sRegions struct {
}

func init() {
	service.RegisterRegions(NewRegions())
}

func NewRegions() *sRegions {
	return &sRegions{}
}

// 增加地区
func (s *sRegions) AddRegions(ctx context.Context, in *admin.AddRegionsReq) (err error) {
	_, err = dao.ReRegions.Ctx(ctx).Insert(in)
	return
}

// 更新地区
func (s *sRegions) UpdateRegions(ctx context.Context, in *admin.UpdateRegionsReq) (err error) {
	affected, err := dao.ReRegions.Ctx(ctx).OmitEmpty().Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该问题不存在")
	}
	return
}

// 删除地区
func (s *sRegions) DeleteRegions(ctx context.Context, in *admin.DeleteRegionsReq) (err error) {
	_, err = dao.ReRegions.Ctx(ctx).Delete("id", in.ID)
	return
}

// 获取地区
func (s *sRegions) GetRegions(ctx context.Context, in *admin.GetRegionsReq) (res *admin.RegionsRes, err error) {
	err = dao.ReRegions.Ctx(ctx).Where("id", in.ID).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 获取地区列表
func (s *sRegions) GetRegionsList(ctx context.Context, in *admin.GetRegionsListReq) (res []*admin.RegionsRes, err error) {
	// 第一步：获取平面结构的区域列表
	var flatRegions []*admin.RegionsRes
	err = dao.ReRegions.Ctx(ctx).Scan(&flatRegions)
	if err != nil {
		return nil, err
	}

	// 第二步：将平面列表转换为映射，便于快速查找
	regionMap := make(map[int]*admin.RegionsRes)
	for _, region := range flatRegions {
		regionMap[region.ID] = &admin.RegionsRes{
			ID:       region.ID,
			ParentID: region.ParentID,
			Sort:     region.Sort,
			Regions:  region.Regions,
			Children: []*admin.RegionsRes{},
		}
	}

	// 第三步：通过将子节点分配给各自的父节点来构建树形结构
	var roots []*admin.RegionsRes
	for _, region := range regionMap {
		if region.ParentID == 0 {
			// 根节点
			roots = append(roots, region)
		} else {
			// 非根节点，找到其父节点并将该节点添加到父节点的子节点列表中
			if parent, exists := regionMap[region.ParentID]; exists {
				parent.Children = append(parent.Children, region)
			}
		}
	}

	// 第四步：对每个节点的子节点进行排序
	var sortChildren func(node *admin.RegionsRes)
	sortChildren = func(node *admin.RegionsRes) {
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].Sort < node.Children[j].Sort
		})
		// 递归对子节点进行排序
		for _, child := range node.Children {
			sortChildren(child)
		}
	}

	// 对所有根节点的子树进行排序
	for _, root := range roots {
		sortChildren(root)
	}

	return roots, nil
}
