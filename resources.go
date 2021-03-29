// AUTOGENERATED CODE. DO NOT EDIT.

package gendoc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ3W/bNhB/z19BaC/bikodugDDQLvA7LnF0GZB3O2dls42MYrUSMqNoel/Hyiq+v5a7bhZkJcguvvxjryP39ESfnMfMnQAqajgM+cH95WDgPsioHw3c/74uHr5k/NmfoWJ1NRnML9CCGuqGcxvpdDCFwwthR+HwDXRVHDsWe0VQkkiCd8BcleUgUpTs1SBb1BGXRhKEveGhJCmlbVmdUQkQe4SlC9pZFZlJip2P4BSZJebLo0jGsycJHFXMWPWsGNNVj2+F3zX4XXIr9HRLXLfEbWiwAJVyLEmGwZoK0kIM4cwVjgsXGKfEaU4CVveSwWyZhsbMiZ2UsQR8gVTM+fHinGEsBFG4BvlJxro/cz53vFORrxyr8dBr5sQvQcSVCUIYSk+1SUIYeBaHufZabFnH7ohH48RDCPekw2wYUglk51A7DX2iL3WQbDeiKCxrlLftWoYPXil4If2jRnlfyHzB3hZ0SYkpqLzKrKP2DOw+bA9s8JEa8xvVuXfUh7APXJ/z+KmkBNAJMEnGgLnnwC2JGYabQlT8F2aYgijPVFUzZcFysVeIU0S4EGadraW7amltfgnYbE5jsHNc9nPKEmaei8D5GanJdXkK4PXE11PK/ZsqxV84WWdXXJD1ULBBr/ea+CGP8/PCDegNASo9DBCDteXIIfHQR9FTE6lkF+IGkHcxOEG5FdmmY4qG43RV2Katr2F4JpQTvmuYblU/Hc2s2kZp7MnxDrYq92gqqqyUHgcXu5OdC6my4I8Rm+vL0FvJ/KSOdv/gE5svM9OJae15aWaqqeN8oe+sV7/9VKUt/l585LBAVj/nMaUb4UMCRvsludJ/jzJnyf5k5rktb6fQj55jaxBHqj/Ze82Hn6Gd8zvD6D34nG8vHhwwrJnReOT/g7+jkFpNE5dd6AiwRVMgD44P+WpvAA55fFpMEkubdGIbfVcu9YSSEj5Lk2Ryv4faufBPdjItzZhxb27sOov3MZjvP00FBVe6nxtu/YJI9Le27Oqrff++I2n/74z1tMd+usxwHn17YZqZSnPtBtJoUV/W3++mAhtAtgPWLx4MaT+jRzIkP72qPd9fGFlb8Wg+28Grb+7HVLfxZtjh75R2i2aapNUOR6z4qtTVF8KPs/M7INFpcMrz0ObN+RmkjOKWkTRJGsmVZOANmeToG+nHWSx3hMZjTveTzuJyWsvsEVcTdpqkFadsjpuThV2usJe8UHq3wAAAP//9p3zJMIaAAA=",
	"html.tmpl": "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",
	"markdown.tmpl": "H4sIAAAAAAAA/+RWzW7jNhC+6ymmUg9rL+TcF7YPjZtdFLtpsA56WRQNbY1tATKpipSRQOS7F/yRSFlW4gLpqT5Ymhlpfr75ZsQEHiom2JYVsGLb+ohUEJEzGs0JUHLERSxYGS/nN2QZRUkCj2RTILAd3DIqkAoeNU1F6B5hdpcXyJWKmubnXV7gX/p1+LSA2T05olIp/Ggad//nh6S7n0QATZNCvoPZN+Sc7JGDUkbrPLdqpQCsm6+M7kNXd3VRhO6QZs5FCkgzSDtJh/mV1sfzGEb3bgGeBVKeMzqI0hlcKA1aWuAJC/A2E9KDqFSKne2a8GusTvl2AGOr9lX+2wpbbQo/1ltSkAr+IEWN8PhSos6aG2V60spUaOUkupognnJdZo548xJIke/pIq7y/UHEyzmBQ4W7RZwYdj6yUj83vyktSbv3o6aZrZBvq7zUpFYqyMaTqhfYw+BZbzx6Rlz0qnH/QvhdjkWmfUowtyANOCDhK9lgARKCN0FGElL9A3uFvuh+IEMMtX/Xb+krBdnRVscLm2rliX3aZGEeNyl/yGmGzzD73eTDIc6wrHBLBGaxzHBH6kLAjhQcJ0pNp6vOOptOW440DWWbCvqQWEBW1oMhiVLgxE9gAAxNzpWp1FKsvXpow9mJpB8XD/EvhOvLfX3cYDUG9RDuDu1x2IOBvgR9D3kr6/1IcprT/bnFpuek/xg6a5r/lNoZPrb7NU2XwSi43fcOcyBBm97qwRWAm+JGwX4LwgCMIQhotr9D4BK3AhhGlnAIycUN/v/l59OAoE+jDPU96bXgjJ3+s3UdQV8h5zcUB5a1HP2Of9fIRdud78hLRjm28mh3zhtxLp7LMvzo6ATG97dL6XyNO3WwzW0TnH4tKiTHnO6VAm7uO8idV1vZ0K3VX/BrDa85fmvwBj3m7anEtdfNZ5LA8CyhmzUr9cm07cY9E8hBwu3HjyDhN3IiIOHhRRzMgH1m2pRo1ZcH3ct68zLWNHu9OFv+z9sDIpo0ffNCNppTtMUwhpsl9FWu07qEbv7KMrTpgkLZVhZqPvd83a4PpCq7pw89Z7r6Vu6A/icAAP//o9X/quoLAAA=",
	"restructuredtext.tmpl": "H4sIAAAAAAAA/8RWzXKbMBC+8xRb0kPr1DwA0/TQOD/TSdJM3OnVls1iMyMERcKNh/LunZUEFhgnadMmXIBdaffT9+1KOhp8vNsiU9ky4zDJlmWKQjGVZMI7MNqrqoKJFUJwnnCUde1V1ds44TgTLEUITyC4YSnWtecFAcwKjGdVZU0hzW7cJ89+dLQJymWR5ASZcrborlFKttIAXSDnJecumKtMrCyg8bOfYUBJDMElk+cJ8qiBs5SbsWILjmEIHRSgtjlCrMd6ABCukUVYhODr+f4H8L9tc6T3FVsgpw8no6/n/EwitZYhsFJlrl4UFOqaxvitEBQhLDAO5w5FlKOu5+QjeJTJjNTLeZeICO8h+KpzSvAjzAtcMoWR/yvCmJVcQcy4xPd1PRpNWm8wGkFVoYjquqpEtiigS5cha2IifGe8JEbsryaq67KhfK+qxqA/PWvyHOLP7hUKSUCfRj6247sCtHFcET4zqd83ZbrA4g/U2KEaVMTiMjpYw2kmFEtEIlZ7LpO++f2/zNoP+PhmrH2Q2laD8fiT04Fnokxfvf16chOmnuYbWnFPanL9laqavoOKPqqSy/c+z0iENiQfKm7L9G5LruvZrqJDz6Nte8xxgxx2k/8B+30t9hvNEh4PAniRVnv9Tpvvtdr8cK/tpO8o3euzKRabZPnoSfeSbbYTWxpwkKJaZ/0j7Voboem2O/xRolTQCH6HMs+ExNbwRL1N2EGxOwedzeecd0Y7a5+qAlmaiBUtQ383Su0HMkAHIhnHg6Ee3Qn2qsGS2haC3TAa8eWScTajs4z6fUp/Behq00zKgauXW056fL+a9CXRrM2Uk2N4uKZM8WUKDxy/TiDgTKxKtkIo6Mog27top2qCnGa0RXF6fEyvL2zD6H27VWuzb1xk2n2krZe3uqLKxXagcpo6caA0G0Ged/4pS8dg0nVMF70Q0zUr8t2EdTciQbKGgS1fGjFI6N8BAAD//0gjAb++CwAA",
	"scalars.json": "H4sIAAAAAAAA/9yXzW4aMRDH7zzFiFMqBZDSlEa9JZGQOOQEOUWp5GVnvW6NTewxzaqq1HfoG/ZJqt0F1gYvpChVk9zQfHg9v/nPWNx1AL53AAC6C6NJT4sFdj9BN9Uukdg9rV1KE9rSvDbMFot45MzG7XxzciY1o+H52vGFLVk8ZZEvgpyNvaBcq6jLuKRYO0aVowPw47SlxiB1X4lBYFNhYN4q8P1ZrMCwjn9dn1DeNTb13Vq0sGRGsERiT6LilAOqmU6F4n0YK8wyMROoCDJtNh5QyBmJJYJy8wSNhd8/f4HIoNDOQCZQpiAsSPEVZQGkIWdl7DppyaRDewrOItjqYiCUJWRpPwI8uHkDXKgI7iDWh+1Fe6iFIuRo4rC9FB/1leDKzUEbGInH8tcJs2DwwQmD6btDPWi0/sJ6MDw/0IPm5k0PpFY82oT4SPvhu10YWDLCC9huxsBP3+3IXvLuGPlHOLgWMbqoGt0zy/EAgaM06Y4RZQuZmERcXCPu/4jkKET2qN05EVxhCkJRPWt9mOZoEebaIGxGWhZ1Cu6OM+VMgUHuJDNQXcG+7fVoj9qPz855eN7G+VWvwEw8YhqR8aX8xgoLWflqJAWh7cNNgK4G5FbvdLZ6OYAZBJ0RKuAGGaGp484+n128vM35zEqtWEakumKJgud0CGb54B6G+WH4epft/mF/uh7f9tL7Cy29wa2UaC2f8q/Lj2vK9K08eqZfZWlHpiKFbnuCCrecfoFT4/BaMmsHIyZt/XN/twOSTa+hdsDcWQJWd36mFTGh4HY66l2sXq+01NjHXiIILifX4zEQPlJMF+GHGmKhnbdczKc2CZuft3wiZGbJDJwS5ZVj3Ooz4aQqbf98VMrfAXbDig0fpgpgJhFkmCnA4oNDNSvXafvUtNG5KggnLYTu7svjYoR2s55OaR+dqsO9i6vxtEbUue/8CQAA//8z+wC/ohEAAA==",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
